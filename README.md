# mutexer

[![travis Status](https://travis-ci.org/mh-cbon/mutexer.svg?branch=master)](https://travis-ci.org/mh-cbon/mutexer) [![Appveyor Status](https://ci.appveyor.com/api/projects/status/github/mh-cbon/mutexer?branch=master&svg=true)](https://ci.appveyor.com/projects/mh-cbon/mutexer) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/mutexer)](https://goreportcard.com/report/github.com/mh-cbon/mutexer) [![GoDoc](https://godoc.org/github.com/mh-cbon/mutexer?status.svg)](http://godoc.org/github.com/mh-cbon/mutexer) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Package mutexer generic sync version of a type using sync.Lock.


# TOC
- [Install](#install)
- [Cli](#cli)
  - [Usage](#usage)
    - [$ mutexer -help](#-mutexer--help)
  - [Cli examples](#cli-examples)
- [API example](#api-example)
  - [> demo/main.go](#-demomaingo)
  - [> demo/tomatessync.go](#-demotomatessyncgo)
- [Recipes](#recipes)
  - [Release the project](#release-the-project)
- [History](#history)

# Install
```sh
mkdir -p $GOPATH/src/github.com/mh-cbon/mutexer
cd $GOPATH/src/github.com/mh-cbon/mutexer
git clone https://github.com/mh-cbon/mutexer.git .
glide install
go install
```

# Cli

## Usage

#### $ mutexer -help
```sh
mutexer 0.0.0

Usage

  mutexer [-p name] [...types]

  types:  A list of types such as src:dst.
          A type is defined by its package path and its type name,
          [pkgpath/]name
          If the Package path is empty, it is set to the package name being generated.
          Name can be a valid type identifier such as TypeName, *TypeName, []TypeName 
  -p:     The name of the package output.
```

## Cli examples

```sh
# Create a mutexed version of demo/Tomate to gen_test/TomateSync
mutexer demo/Tomate:gen_test/TomateSync
# Specify the out pkg name
mutexer -p demo demo/Tomate:gen_test/TomateSync
# Use stdout
mutexer -p demo - demo/Tomate:gen_test/TomateSync
```
# API example

Following example demonstates a program using it to generate a sync version of a type.

#### > demo/main.go
```go
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
```

Following code is the generated implementation of `TomateSync` type.

#### > demo/tomatessync.go
```go
package main

// file generated by
// github.com/mh-cbon/mutexer
// do not edit

import (
	"sync"
)

// TomatesSync mutexes a Tomates
type TomatesSync struct {
	embed *Tomates
	mutex *sync.Mutex
}

// NewTomatesSync constructs a new TomatesSync
func NewTomatesSync() *TomatesSync {
	ret := &TomatesSync{}
	embed := NewTomates()
	ret.embed = embed
	ret.mutex = &sync.Mutex{}
	return ret
}

// Push is mutexed
func (t *TomatesSync) Push(x ...Tomate) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Push(x...)
}

// Unshift is mutexed
func (t *TomatesSync) Unshift(x ...Tomate) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Unshift(x...)
}

// Pop is mutexed
func (t *TomatesSync) Pop() Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Pop()
}

// Shift is mutexed
func (t *TomatesSync) Shift() Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Shift()
}

// Index is mutexed
func (t *TomatesSync) Index(s Tomate) int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Index(s)
}

// Contains is mutexed
func (t *TomatesSync) Contains(s Tomate) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Contains(s)
}

// RemoveAt is mutexed
func (t *TomatesSync) RemoveAt(i int) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.RemoveAt(i)
}

// Remove is mutexed
func (t *TomatesSync) Remove(s Tomate) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Remove(s)
}

// InsertAt is mutexed
func (t *TomatesSync) InsertAt(i int, s Tomate) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.InsertAt(i, s)
}

// Splice is mutexed
func (t *TomatesSync) Splice(start int, length int, s ...Tomate) []Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Splice(start, length, s...)
}

// Slice is mutexed
func (t *TomatesSync) Slice(start int, length int) []Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Slice(start, length)
}

// Reverse is mutexed
func (t *TomatesSync) Reverse() *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Reverse()
}

// Len is mutexed
func (t *TomatesSync) Len() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Len()
}

// Set is mutexed
func (t *TomatesSync) Set(x []Tomate) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Set(x)
}

// Get is mutexed
func (t *TomatesSync) Get() []Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Get()
}

// At is mutexed
func (t *TomatesSync) At(i int) Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.At(i)
}

// Filter is mutexed
func (t *TomatesSync) Filter(filters ...func(Tomate) bool) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Filter(filters...)
}

// Map is mutexed
func (t *TomatesSync) Map(mappers ...func(Tomate) Tomate) *Tomates {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Map(mappers...)
}

// First is mutexed
func (t *TomatesSync) First() Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.First()
}

// Last is mutexed
func (t *TomatesSync) Last() Tomate {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Last()
}

// Empty is mutexed
func (t *TomatesSync) Empty() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.embed.Empty()
}
```


# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)
