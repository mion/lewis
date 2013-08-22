package main

import (
	"fmt"
	. "lewis"
)

func main() {
	in := "(quote 32)"
	p := Parse(in)
	fmt.Println(in)
	fmt.Println(Eval(p, NewScope(nil)))
	// var a Any
	// s := "test"
	// a = NewSymbol(s)
	// fmt.Println(IsSymbol(a))
	// fmt.Println(IsSymbol(s))
	// fmt.Println(IsSymbol("Lisp"))
}
