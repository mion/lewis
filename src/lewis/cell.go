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
	} else if c == nil {
		return ""
	} else {
		return fmt.Sprintf(" %v", c.Car) + c.Cdr.toString(false)
	}
}

func (c *Cell) String() string {
	return c.toString(true)
}

func (c *Cell) Push(a Any) {
	if c.Cdr != nil {
		c.Cdr.Push(a)
	} else {
		c.Cdr = NewCell(a)
	}
}

func (c *Cell) Pop() Any {
	return nil
}
