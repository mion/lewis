package lewis

import (
	"fmt"
)

type Cell struct {
	Car Any
	Cdr *Cell
}

func Cons(car Any, cdr *Cell) *Cell {
	if cdr != nil && cdr.Car == nil && cdr.Cdr == nil {
		cdr.Car = car
		return cdr
	} else {
		return &Cell{car, cdr}
	}
}

func NewCell(v Any) *Cell {
	l := new(Cell)
	l.Car = v
	return l
}

func (c *Cell) toString(root bool) string {
	if root {
		return "(" + fmt.Sprintf("%v", c.Car) + c.Cdr.toString(false) + ")"
	} else if c == nil || c.IsNull() {
		return ""
	} else {
		return fmt.Sprintf(" %v", c.Car) + c.Cdr.toString(false)
	}
}

func (c *Cell) String() string {
	return c.toString(true)
}

func (c *Cell) Print() {
	if c == nil {
		fmt.Print("nil")
	} else if c.IsNull() {
		fmt.Print("~")
	} else {
		fmt.Print("( ")
		if cell, ok := c.Car.(*Cell); ok {
			cell.Print()
		} else {
			fmt.Printf("*")
		}
		fmt.Print(" <> ")
		c.Cdr.Print()
		fmt.Print(" )")
	}
}

func ident(str string, n int) string {
	tabs := ""
	for i := 0; i < n; i++ {
		tabs += "\t"
	}
	return tabs + str + "\n"
}

func inspect(x Any, k int) string {
	if x == nil {
		return ident("nil", k)
	}
	if c, ok := x.(*Cell); ok {
		if c.IsNull() {
			return ident("~", k)
		} else {
			n := k
			if _, ok := c.Car.(*Cell); ok {
				n = k + 1
			}
			return inspect(c.Car, n) + inspect(c.Cdr, k)
		}
	} else if a, ok := x.(*Atom); ok {
		return ident(fmt.Sprintf("%v", a.Value), k)
	} else {
		return ident(fmt.Sprintf("%v", x), k) // in case something weird is in there
	}
}

func (c *Cell) Inspect() {
	fmt.Println(inspect(c, 0))
}

func (c *Cell) IsNull() bool {
	return c.Car == nil && c.Cdr == nil
}

func (c *Cell) Cadr(n int) Any {
	p := c
	for n > 0 {
		p = p.Cdr
		n--
	}
	return p.Car
}

func (c *Cell) Len() int {
	if c == nil || c.IsNull() {
		return 0
	} else {
		return 1 + c.Cdr.Len()
	}
}
