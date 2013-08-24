package main

import (
	"fmt"
	"math/big"
	r "reflect"
)

func Analyze(a interface{}) {
	fmt.Println("------")

	t := r.TypeOf(a)
	v := r.ValueOf(a)
	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Size:", t.Size())
	fmt.Println("Methods:", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i))
	}
}

func main() {
	a := 42
	b := 3.4
	c := "Hello"
	d := big.NewInt(0)
	Analyze(a)
	Analyze(b)
	Analyze(c)
	Analyze(d)
}
