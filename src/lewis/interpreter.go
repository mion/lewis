package lewis

import (
	"fmt"
)

func Eval(x Any, s *Scope) Any {
	Debug("Evaluating", x, "in scope", s)
	if sym, ok := x.(*Symbol); ok {
		if scope := s.Find(sym); scope != nil {
			Debug("Found symbol", x, "defined in scope", scope)
			return scope.Get(sym) // variable reference
		} else {
			Debug("No symbol", x, "found in scope", scope)
			return nil
		}
	} else if c, ok := x.(*Cell); ok {
		if fn, ok := Eval(c.Car, s).(func(*Cell, *Scope) Any); ok {
			Debug("Found function", fn, "in cell", c, "and scope", s)
			return fn(c, s)
		} else {
			Debug("Expected a function instead of", c.Car)
			return Error(c.Car, "is not a function")
		}
	} else {
		Debug("Constant literal")
		return x // literal
	}
}

func ParseEval(in string, s *Scope) Any {
	return Eval(Parse(in), s)
}

func ParseEvalPrint(in string, s *Scope) {
	fmt.Println(ParseEval(in, s))
}
