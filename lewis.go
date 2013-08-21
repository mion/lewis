package main

import (
	"fmt"
	"interp"
	r "reflect"
)

func main() {
	a := interp.NewEnv(nil)
	b := interp.NewEnv(a)
	a.Set("x", r.ValueOf(5))
	if e := b.Find("x"); e != nil {
		x := e.Get("x")
		fmt.Println(x.Interface())
		fmt.Println(x.Kind() == r.Int)
	}
	fmt.Println(interp.Eval(r.ValueOf("x"), a).Interface())
}

// package main

// import (
// 	"bufio"
// 	"container/list"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// type Vector struct {
// 	l *list.List
// }

// func New() *Vector {
// 	a := new(Vector)
// 	a.l = list.New()
// 	return a
// }

// func (a *Vector) Shift() interface{} {
// 	if a.Len() == 0 {
// 		return nil
// 	} else {
// 		return a.l.Remove(a.l.Front())
// 	}
// }

// func (a *Vector) Push(v interface{}) {
// 	a.l.PushBack(v)
// }

// func (a *Vector) Pop() interface{} {
// 	return a.l.Remove(a.l.Back())
// }

// func (a *Vector) Get(index int) interface{} {
// 	i := 0
// 	for e := a.l.Front(); e != nil; e = e.Next() {
// 		if i == index {
// 			return e.Value
// 		}
// 		i++
// 	}
// 	return nil
// }

// func (a *Vector) Set(index int, v interface{}) {
// 	i := 0
// 	for e := a.l.Front(); e != nil; e = e.Next() {
// 		if i == index {
// 			e.Value = v
// 			return
// 		}
// 		i++
// 	}
// }

// func (a *Vector) ToSlice() []interface{} {
// 	s := make([]interface{}, a.l.Len())
// 	i := 0
// 	for e := a.l.Front(); e != nil; e = e.Next() {
// 		s[i] = e.Value
// 		i++
// 	}
// 	return s
// }

// func (a *Vector) FromSlice(s []interface{}) {
// 	for i := range s {
// 		a.l.PushBack(s[i])
// 	}
// }

// func (a *Vector) Len() int {
// 	return a.l.Len()
// }

// func (a *Vector) Concat(v interface{}) *Vector {
// 	b := New()
// 	b.FromSlice(a.ToSlice())
// 	b.Push(v)
// 	return b
// }

// type Object interface{}

// func tokenize(s string) *Vector {
// 	var leftParens = regexp.MustCompile(`\(`)
// 	var rightParens = regexp.MustCompile(`\)`)

// 	s = leftParens.ReplaceAllString(s, " ( ")
// 	s = rightParens.ReplaceAllString(s, " ) ")
// 	s = strings.TrimSpace(s)
// 	fields := strings.Fields(s)

// 	a := New()
// 	n := make([]interface{}, len(fields))
// 	for i, v := range fields {
// 		n[i] = interface{}(v)
// 	}
// 	a.FromSlice(n)

// 	return a
// }

// func isInteger(input string) bool {
// 	var r = regexp.MustCompile(`^[0-9]+$`)
// 	m := r.MatchString(input)
// 	return m
// }

// func isString(input string) bool {
// 	return strings.HasPrefix(input, "\"") && strings.HasSuffix(input, "\"")
// }

// func categorize(input string) map[string]Object {
// 	m := make(map[string]Object)

// 	// Is it an integer?
// 	if isInteger(input) {
// 		i, err := strconv.ParseInt(input, 0, 32)
// 		if err == nil {
// 			m["type"] = "literal"
// 			m["value"] = i
// 		} else {
// 			panic("unable to parse integer")
// 		}
// 	} else if isString(input) {
// 		m["type"] = "literal"
// 		m["value"] = strings.TrimSuffix(strings.TrimPrefix(input, "\""), "\"")
// 	} else {
// 		m["type"] = "identifier"
// 		m["value"] = input
// 	}

// 	return m
// }

// func parenthesize(input *Vector, l *Vector) interface{} {
// 	if l == nil {
// 		return parenthesize(input, New())
// 	} else {
// 		token := input.Shift()
// 		if token == nil {
// 			return l.Pop()
// 		} else if token == "(" {
// 			l.Push(parenthesize(input, New()))
// 			return parenthesize(input, l)
// 		} else if token == ")" {
// 			return l
// 		} else {
// 			str := token.(string)
// 			return parenthesize(input, l.Concat(categorize(str)))
// 		}
// 	}
// }

// func repl() {
// 	reader := bufio.NewReader(os.Stdin)
// 	var cmd string

// 	for cmd != "exit" {
// 		fmt.Print(">>> ")
// 		input, err := reader.ReadString('\n')
// 		if err == nil {
// 			cmd = strings.TrimSpace(input)
// 			fmt.Println(cmd)
// 		}
// 	}
// }

// func main() {
// 	s := "((lambda (x) x) \"Lisp\")"
// 	fmt.Println(tokenize(s).ToSlice())
// 	fmt.Println(categorize("31"))
// 	fmt.Println(categorize("def"))
// 	fmt.Println(categorize("\"foo\""))
// 	p := parenthesize(tokenize(s), nil)
// 	v := p.(*Vector)
// 	fmt.Println(v.ToSlice())
// }
