package main

import (
	"fmt"
	. "lewis"
)

func main() {
	fmt.Println(Parse("(lambda (x) (* x x))"))
	fmt.Println(Parse("(a b (c d) e)"))
}
