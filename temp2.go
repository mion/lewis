package main

import (
	"fmt"
	r "reflect"
)

type Any interface{}

type Cell struct {
	Car Any
	Cdr *Cell
}

func Cons(car Any, cdr *Cell) *Cell {
	return &Cell{car, cdr}
}

func (c *Cell) Cons(car Any) *Cell {
	return &Cell{car, c}
}

func NewCell(v Any) *Cell {
	l := new(Cell)
	l.Car = v
	return l
}

func (c *Cell) toString(root bool) string {
	if root {
		return "(" + fmt.Sprintf("%v", c.Car) + c.Cdr.toString(false) + ")"
	} else if c == nil {
		return ""
	} else {
		return fmt.Sprintf(" %v", c.Car) + c.Cdr.toString(false)
	}
}

func (c *Cell) String() string {
	return c.toString(true)
}

func (c *Cell) IsLeaf() bool {
	return c.Cdr == nil
}

func (c *Cell) IsList() bool {
	switch c.Car.(type) {
	case *Cell:
		return true
	default:
		return false
	}
}

type Symbol struct {
	string
}

func eval(x interface{}) interface{} {
	switch t := x.(type) {
	case string:
		return t + t
	case int:
		return t * t
	case *Cell:
		return t.Car
	default:
		return nil
	}
}

func find(name Symbol) string {
	return name.String()
}

func main() {
	// (lambda (x) (* x x))
	l := NewCell("x")
	l = l.Cons("x")
	l = l.Cons("*")

	l = NewCell(l)
	l = l.Cons(NewCell("x"))

	l = l.Cons("lambda")

	// fmt.Println(args.String())
	// fmt.Println(exp.String())
	// x := r.ValueOf(l)
	fmt.Println(r.ValueOf(l))
	// if cell, ok := x.Interface().(*Cell); ok {
	// 	fmt.Println(cell.Car)
	// }
	x := "123"
	y := 12
	z := l
	var i Any
	var j Any
	i = x
	j = Symbol{"321"}
	fmt.Println(eval(x))
	fmt.Println(eval(y))
	fmt.Println(eval(z))
	fmt.Println(i == "123")
	if sym, ok := j.(Symbol); ok {
		fmt.Println(find(sym))
	}
}
