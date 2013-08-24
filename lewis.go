package main

import (
	. "lewis"
)

func main() {
	s := NewScope(GlobalScope)
	ParseEvalPrint("(define add (lambda (a b) (+ a b)))", s)
	ParseEvalPrint("(add 5 3)", s)
	// examples := [...]string{
	// 	"12",
	// 	"(quote (a b c))",
	// 	"(if (< 10 20) (+ 1 1) (+ 3 3))",
	// 	"(define x 2)",
	// 	"(+ x 1)",
	// 	"(set! x (* x x))",
	// 	"(define r 3)",
	// 	"(lambda (r) (* 3.1415 (* r r)))",
	// 	"(begin (set! x 1) (set! x (+ x 1)) (* x 2))",
	// 	"(define square (lambda (x) (* x x))) (square 12)",
	// }
}
