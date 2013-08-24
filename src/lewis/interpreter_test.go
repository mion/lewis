package lewis

import (
	"fmt"
	"testing"
)

func checkInt64(t *testing.T, x Any, n int64) bool {
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

func TestConstantLiteral(t *testing.T) {
	var x Any
	s := NewScope(GlobalScope)

	x = Eval(Parse("12"), s)
	checkInt64(t, x, 12)
	x = Eval(Parse("-1"), s)
	checkInt64(t, x, -1)
	x = Eval(Parse("\"Lisp\""), s)
	checkString(t, x, "Lisp")
}

func TestScope(t *testing.T) {
	x := Sym("x")
	s := NewScope(GlobalScope)
	s.table[x] = int64(42)

	checkInt64(t, s.Get(x), 42)
}

func checkScope(t *testing.T, sym *Symbol, start *Scope, expected *Scope) bool {
	if s := start.Find(sym); s != expected {
		t.Errorf("variable %s found in scope %v, expected %v", sym, s, expected)
		return false
	} else {
		return true
	}
}

func TestDefine(t *testing.T) {
	x := Sym("x")
	s1 := NewScope(GlobalScope)
	s2 := NewScope(s1)
	s3 := NewScope(s2)
	res := Eval(Parse("(define x 2)"), s3)
	checkNil(t, res)
	checkScope(t, x, s3, s3)
	checkInt64(t, s3.Get(x), 2)
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
		{in: "(define x 7)", out: "<nil>"},
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
