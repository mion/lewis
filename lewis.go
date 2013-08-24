package main

import (
	"fmt"
	. "lewis"
)

func main() {
	examples := [...]string{
		"12",
		"(quote (a b c))",
		"(if (< 10 20) (+ 1 1) (+ 3 3))",
		"(define x 2)",
		"(+ x 1)",
		"(set! x (* x x))",
		"(define r 3)",
		"(lambda (r) (* 3.1415 (* r r)))",
		"(begin (set! x 1) (set! x (+ x 1)) (* x 2))",
		"(define square (lambda (x) (* x x))) (square 12)",
	}

	in := examples[3]
	fmt.Println(in)
	p := Parse(in)
	fmt.Println(Eval(p, GlobalScope))
	fmt.Println(Eval("x", GlobalScope))
	in = examples[4]
	fmt.Println(in)
	p = Parse(in)
	fmt.Println(Eval(p, GlobalScope))

	// var a Any
	// s := "test"
	// a = NewSymbol(s)
	// fmt.Println(IsSymbol(a))
	// fmt.Println(IsSymbol(s))
	// fmt.Println(IsSymbol("Lisp"))
}
