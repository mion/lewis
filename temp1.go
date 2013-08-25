package main

import (
	"fmt"
)

type Lib struct {
	name string
}

func main() {
	n := make([]*Lib, 0)
	n = append(n, &Lib{name: "hello"})
	fmt.Println(n)
}
