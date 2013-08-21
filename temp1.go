package main

import (
	"fmt"
)

type Scope struct {
	integers map[string]int
}

func (s *Scope) Init() *Scope {
	s.integers = make(map[string]int)
	return s
}

func NewScope() *Scope {
	return new(Scope).Init()
}

func main() {
	s := NewScope()
	s.integers["x"] = 12
	fmt.Println(s.integers["x"])
}
