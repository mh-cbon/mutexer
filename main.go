// Package mutexer generic sync version of a type using sync.Lock.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mh-cbon/astutil"
	"github.com/mh-cbon/lister/utils"
)

var name = "mutexer"
var version = "0.0.0"

func main() {

	var help bool
	var h bool
	var ver bool
	var v bool
	var outPkg string
	flag.BoolVar(&help, "help", false, "Show help.")
	flag.BoolVar(&h, "h", false, "Show help.")
	flag.BoolVar(&ver, "version", false, "Show version.")
	flag.BoolVar(&v, "v", false, "Show version.")
	flag.StringVar(&outPkg, "p", os.Getenv("GOPACKAGE"), "Package name of the new code.")
	flag.Parse()

	if ver || v {
		showVer()
		return
	}
	if help || h {
		showHelp()
		return
	}

	if flag.NArg() < 1 {
		panic("wrong usage")
	}
	args := flag.Args()

	out := ""
	if args[0] == "-" {
		args = args[1:]
		out = "-"
	}

	todos, err := utils.NewTransformsArgs(utils.GetPkgToLoad()).Parse(args)
	if err != nil {
		panic(err)
	}

	filesOut := utils.NewFilesOut("github.com/mh-cbon/" + name)

	for _, todo := range todos.Args {
		srcName := todo.FromTypeName
		destName := todo.ToTypeName
		toImport := todo.FromPkgPath
		if todo.FromPkgPath == "" {
			log.Println("Skipped ", srcName)
			continue
		}
		prog := astutil.GetProgramFast(toImport).Package(toImport)
		foundTypes := astutil.FindTypes(prog)
		foundMethods := astutil.FindMethods(prog)
		foundCtors := astutil.FindCtors(prog, foundTypes)

		fileOut := filesOut.Get(todo.ToPath)
		fileOut.PkgName = outPkg

		if fileOut.PkgName == "" {
			fileOut.PkgName = findOutPkg(todo)
		}
		fileOut.AddImport("sync", "")

		if todo.FromPkgPath != todo.ToPkgPath {
			fileOut.AddImport(todo.FromPkgPath, "")
		}
		if todo.FromPkgPath != todo.ToPkgPath {
			fileOut.AddImport(todo.FromPkgPath, "")
		}

		res := processType(destName, srcName, foundCtors, foundMethods)
		io.Copy(&fileOut.Body, &res)
	}

	filesOut.Write(out)
}

func showVer() {
	fmt.Printf("%v %v\n", name, version)
}

func showHelp() {
	showVer()
	fmt.Println()
	fmt.Println("Usage")
	fmt.Println()
	fmt.Printf("  %v [-p name] [...types]\n\n", name)
	fmt.Printf("  types:  A list of types such as src:dst.\n")
	fmt.Printf("          A type is defined by its package path and its type name,\n")
	fmt.Printf("          [pkgpath/]name\n")
	fmt.Printf("          If the Package path is empty, it is set to the package name being generated.\n")
	// fmt.Printf("          If the Package path is a directory relative to the cwd, and the Package name is not provided\n")
	// fmt.Printf("          the package path is set to this relative directory,\n")
	// fmt.Printf("          the package name is set to the name of this directory.\n")
	fmt.Printf("          Name can be a valid type identifier such as TypeName, *TypeName, []TypeName \n")
	fmt.Printf("  -p:     The name of the package output.\n")
	fmt.Println()
}

func findOutPkg(todo utils.TransformArg) string {
	if todo.ToPkgPath != "" {
		prog := astutil.GetProgramFast(todo.ToPkgPath)
		if prog != nil {
			pkg := prog.Package(todo.ToPkgPath)
			return pkg.Pkg.Name()
		}
	}
	if todo.ToPkgPath == "" {
		prog := astutil.GetProgramFast(utils.GetPkgToLoad())
		if len(prog.Imported) < 1 {
			panic("impossible, add [-p name] option")
		}
		for _, p := range prog.Imported {
			return p.Pkg.Name()
		}
	}
	if strings.Index(todo.ToPkgPath, "/") > -1 {
		return filepath.Base(todo.ToPkgPath)
	}
	return todo.ToPkgPath
}

func processType(destName, srcName string, foundCtors map[string]*ast.FuncDecl, foundMethods map[string][]*ast.FuncDecl) bytes.Buffer {

	var b bytes.Buffer
	dest := &b

	srcConcrete := astutil.GetUnpointedType(srcName)
	dstConcrete := astutil.GetUnpointedType(destName)

	fmt.Fprintf(dest, "// %v mutexes a %v\n", dstConcrete, srcConcrete)
	fmt.Fprintf(dest, `type %v struct{
		 embed %v
		 mutex *sync.Mutex
		 }
		 `, dstConcrete, srcName)

	ctorParams := ""
	ctorName := ""
	ctorIsPointer := false
	if x, ok := foundCtors[srcConcrete]; ok {
		ctorParams = astutil.MethodParams(x)
		ctorIsPointer = astutil.MethodReturnPointer(x)
		ctorName = "New" + srcConcrete
	}

	fmt.Fprintf(dest, `// New%v constructs a new %v
		`, dstConcrete, destName)
	fmt.Fprintf(dest, `func New%v(%v) *%v {
		`, dstConcrete, ctorParams, dstConcrete)
	fmt.Fprintf(dest, `ret := &%v{}
		`, dstConcrete)
	if ctorName != "" {
		fmt.Fprintf(dest, "	embed := %v(%v)\n", ctorName, ctorParams)
		if !ctorIsPointer && astutil.IsAPointedType(srcName) {
			fmt.Fprintf(dest, "	ret.embed = *embed\n")
		} else {
			fmt.Fprintf(dest, "	ret.embed = embed\n")
		}
	}
	fmt.Fprintf(dest, `	ret.mutex = &sync.Mutex{}
		return ret
		}
	`)

	for _, m := range foundMethods[srcConcrete] {
		withEllipse := astutil.MethodHasEllipse(m)
		paramNames := astutil.MethodParamNamesInvokation(m, withEllipse)
		receiverName := astutil.ReceiverName(m)
		methodName := astutil.MethodName(m)
		callExpr := fmt.Sprintf("%v.embed.%v(%v)", receiverName, methodName, paramNames)
		sExpr := fmt.Sprintf(`
  %v.mutex.Lock()
  defer %v.mutex.Unlock()
	return %v`,
			receiverName, receiverName, callExpr)
		sExpr = fmt.Sprintf("func(){%v\n}", sExpr)
		expr, err := parser.ParseExpr(sExpr)
		if err != nil {
			panic(err)
		}
		astutil.SetReceiverTypeName(m, destName)
		astutil.SetReceiverPointer(m, true)
		m.Body = expr.(*ast.FuncLit).Body
		m.Doc = nil // clear the doc.
		fmt.Fprintf(dest, "// %v is mutexed\n", methodName)
		fmt.Fprintf(dest, "%v\n", astutil.Print(m))
	}

	return b
}
