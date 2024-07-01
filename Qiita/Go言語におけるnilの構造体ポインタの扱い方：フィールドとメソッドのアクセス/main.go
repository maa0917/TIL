package main

import "fmt"

type Foo struct{}

func (f *Foo) Bar() string {
	return "bar"
}

func main() {
	var foo *Foo
	if foo != nil {
		fmt.Println(foo.Bar())
	} else {
		fmt.Println("foo is nil")
	}
}
