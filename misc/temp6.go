package main

import (
	"fmt"
	. "lewis"
)

func main() {
	a := "()"
	b := "5"
	c := "(3)"
	d := "(x)"
	e := "(x y z)"
	fmt.Println(Parse(a))
	fmt.Println(Parse(b))
	fmt.Println(Parse(c))
	fmt.Println(Parse(d))
	fmt.Println(Parse(e))
	fmt.Println("\n\n")

	in := "((lambda (x) x) \"Lisp\")"
	p := Parse(in).(*Cell)
	fmt.Println(in + "\n")
	fmt.Println(p)
	fmt.Println("\n")
	p.Inspect()
	fmt.Println("\n")
	p.Print()
	fmt.Println("\n---\n")

	Parse(a).(*Cell).Print()
	fmt.Println("\n")
	Parse(c).(*Cell).Print()
	fmt.Println("\n")
	Parse(d).(*Cell).Print()
	fmt.Println("\n")
	Parse(e).(*Cell).Print()
	fmt.Println("\n")
}
