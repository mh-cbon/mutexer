package main

import "fmt"

//go:generate lister Tomate:Tomates
//go:generate mutexer *Tomates:TomatesSync

// Tomate is about red vegetables to make famous italian food.
type Tomate struct {
	name string
}

// GetID return the ID of the Tomate.
func (t Tomate) GetID() string {
	return t.name
}

// Hello world!
func (t *Tomate) Hello() { fmt.Println(" world!") }

// Good bye!
func (t Tomate) Good() { fmt.Println(" bye!") }

// Name it!
func (t Tomate) Name(it string) string { return fmt.Sprintf("Hello %v!\n", it) }

// NewTomate is a contrstuctor
func NewTomate(n string) *Tomate {
	return &Tomate{
		name: n,
	}
}

func main() {
	slice := NewTomatesSync()
	slice.Push(Tomate{"Red"})
	fmt.Println(
		slice.Filter(FilterTomates.Byname("Red")).First().Name("world"),
	)
}
