package lewis

import (
	"fmt"
	"testing"
)

func checkInt(t *testing.T, x Any, n int64) bool {
	if i, ok := x.(int64); ok && i == n {
		return true
	} else {
		t.Errorf("x = %v, want %d", x, n)
		return false
	}
}

func checkString(t *testing.T, x Any, s string) bool {
	if q, ok := x.(string); ok && q == s {
		return true
	} else {
		t.Errorf("x = %v, want %s", x, s)
		return false
	}
}

func checkNil(t *testing.T, x Any) bool {
	if x != nil {
		t.Errorf("x = %v, want <nil>", x)
		return false
	} else {
		return true
	}
}

func checkCellLen(t *testing.T, c *Cell, len int) bool {
	if n := c.Len(); n != len {
		t.Errorf("c.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func TestCell(t *testing.T) {
	c := Cons(nil, nil)
	checkCellLen(t, c, 0)
}

func TestConstantLiteral(t *testing.T) {
	var x Any
	s := NewScope(GlobalScope)

	x = Eval(Parse("12"), s)
	checkInt(t, x, 12)
	x = Eval(Parse("-1"), s)
	checkInt(t, x, -1)
	x = Eval(Parse("\"Lisp\""), s)
	checkString(t, x, "Lisp")
}

func TestScope(t *testing.T) {
	x := Sym("x")
	s := NewScope(GlobalScope)
	s.table[x] = int64(42)

	checkInt(t, s.Get(x), 42)
}

func checkDefined(t *testing.T, name string, s *Scope) bool {
	if scp := s.Find(Sym(name)); s != scp {
		t.Errorf("s.Find(%s) = %v, want %v", name, scp, s)
		return false
	} else {
		return true
	}
}

func checkScope(t *testing.T, s *Scope, table map[string]Any) {
	for name, value := range table {
		if !checkDefined(t, name, s) {
			return
		}
		if v := s.Get(Sym(name)); v != ToInt64(value) {
			t.Errorf("s.Get(%s) = %v, want %v", name, v, value)
		}
	}
}

func TestDefineAndSet(t *testing.T) {
	above := NewScope(GlobalScope)
	local := NewScope(above)

	Eval(Parse("(def x 2)"), local)
	checkScope(t, local, map[string]Any{"x": 2})

	Eval(Parse("(def y 3)"), above)
	checkScope(t, local, map[string]Any{"x": 2})
	checkScope(t, above, map[string]Any{"y": 3})

	Eval(Parse("(def y 7)"), local)
	checkScope(t, above, map[string]Any{"y": 3})
	checkScope(t, local, map[string]Any{"x": 2, "y": 7})

	Eval(Parse("(set! y 5)"), local)
	checkScope(t, above, map[string]Any{"y": 3})
	checkScope(t, local, map[string]Any{"x": 2, "y": 5})

	Eval(Parse("(def z 9)"), above)
	Eval(Parse("(set! z -3)"), local)
	checkScope(t, above, map[string]Any{"y": 3, "z": -3})
	checkScope(t, local, map[string]Any{"x": 2, "y": 5})
}

func TestLambda(t *testing.T) {
	s := NewScope(GlobalScope)
	if lambda, ok := ParseEval("(func (x) (* x x))", s).(func(*Cell, *Scope) Any); ok {
		c := Cons(Sym("qwerty"), Cons(int64(4), Cons(nil, nil)))
		r := lambda(c, s)
		checkInt(t, r, 16)
	}
	if lambda, ok := ParseEval("(func (a b c) (+ a (+ b c)))", s).(func(*Cell, *Scope) Any); ok {
		c := Cons(Sym("xyz"), Cons(int64(1), Cons(int64(2), Cons(int64(3), Cons(nil, nil)))))
		r := lambda(c, s)
		checkInt(t, r, 6)
	}
}

func TestBegin(t *testing.T) {
	s := NewScope(GlobalScope)
	res := ParseEval("(do (def x 3) (def y 4) (* x y))", s)
	checkInt(t, res, 12)
	checkScope(t, s, map[string]Any{"x": 3, "y": 4})
}

func TestEval(t *testing.T) {
	s := NewScope(GlobalScope)
	tests := []struct {
		in  string
		out string
	}{
		{in: "9", out: "9"},
		{in: "\"Hello\"", out: "Hello"},
		{in: "(quote (a b c))", out: "(a b c)"},
		{in: "(* 2 3)", out: "6"},
		{in: "(+ 2 3)", out: "5"},
		{in: "(- 2 3)", out: "-1"},
		{in: "(< 5 10)", out: "true"},
		{in: "(if (< 10 20) (+ 1 1) (+ 3 3))", out: "2"},
		{in: "(def x 7)", out: "<nil>"},
		{in: "x", out: "7"},
		{in: "(+ x x)", out: "14"},
	}
	for _, tst := range tests {
		res := Eval(Parse(tst.in), s)
		if fmt.Sprint(res) != tst.out {
			t.Errorf("%s ==> %v; expected %s", tst.in, res, tst.out)
		}
	}
}
