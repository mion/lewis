package main

import (
	"fmt"
	rfl "reflect"
)

type Any interface{}

func Inspect(a Any) {
	fmt.Println("Type:", rfl.TypeOf(a))
}

func Compare(a Any, b Any) {
	aType := rfl.TypeOf(a)
	bType := rfl.TypeOf(b)
}

/*
(func (x) (* x x))
*/

func main() {
	i := 5
	j := int64(3)
	Compare(i, j)
}
