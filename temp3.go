package main

import (
	"fmt"
	. "lewis"
)

func main() {
	args := Cons("x", Cons(nil, nil))
	exp := Cons("*", Cons("x", Cons("x", nil)))
	l := Cons("lambda", Cons(args, Cons(exp, nil)))

	fmt.Println(l)

	a, err := Categorize("32")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
	a, err = Categorize("\"hello\"")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
	a, err = Categorize("lambda")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

	fmt.Println(Tokenize("(lambda (r) (* r r 3.14))"))
}
