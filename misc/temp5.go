package main

import (
	"fmt"
	. "lewis"
)

func main() {
	var a Any
	a = "hey"
	fmt.Println(a == 3)
	fmt.Println(Parse("(lambda (x) (* x x))"))
	fmt.Println(Parse("(a b (c d) e)"))
	fmt.Println(Parse("x"))
	fmt.Println(Parse("(x 32 \"a\")"))
	c := Cons("a", Cons("b", Cons("c", nil)))
	fmt.Println(c)
	fmt.Println(c.Cadr(0))
	fmt.Println(c.Cadr(1))
	fmt.Println(c.Cadr(2))
}
