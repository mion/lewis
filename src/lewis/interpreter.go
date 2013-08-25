package lewis

import (
	"fmt"
	rfl "reflect"
)

func checkType(t rfl.Type, a Any) bool {
	if !rfl.TypeOf(a).AssignableTo(t) {
		Error(rfl.TypeOf(a), " is not assignable to ", t)
		return false
	} else {
		Debug("No convertion needed for argument", a)
		return true
	}
}

// Probably a bug... or lack of understanding. Check reflection.go
func convertTypeIfNecessary(t rfl.Type, a Any) rfl.Value {
	if !rfl.TypeOf(a).AssignableTo(t) {
		if rfl.TypeOf(a).ConvertibleTo(t) {
			Debug("Converting", a, "with type", rfl.TypeOf(a), "to type", t, "result", rfl.ValueOf(a).Convert(t).Interface())
			return rfl.ValueOf(a).Convert(t)
		} else {
			Error(rfl.TypeOf(a), " is not assignable nor convertible to ", t)
			return rfl.ValueOf(nil)
		}
	} else {
		Debug("No convertion needed for argument", a)
		return rfl.ValueOf(a)
	}
}

func inspectValues(a []rfl.Value) []string {
	r := make([]string, len(a))
	for i, _ := range a {
		r[i] = inspectValue(a[i])
	}
	return r
}

func inspectValue(v rfl.Value) string {
	return fmt.Sprintf("<Value= %v, Type= %v, Interfaced= %v>", v, v.Type(), v.Interface())
}

func callGoFunction(fn rfl.Value, c *Cell, s *Scope) Any {
	if fn.Type().NumIn() != c.Len()-1 {
		Error("wrong number of arguments for function", fn.Type().Name(), "; got ", c.Len()-1, "expected", fn.Type().NumIn())
	}
	in := make([]rfl.Value, c.Len()-1)
	for i := 0; i < c.Len()-1; i++ {
		arg := Eval(c.Cadr(i+1), s)
		if checkType(fn.Type().In(i), arg) {
			in[i] = rfl.ValueOf(arg)
		}
		// in[i] = convertTypeIfNecessary(fn.Type().In(i), c.Cadr(i+1))
	}
	Debug("Input array (after convertion)", inspectValues(in))
	res := fn.Call(in)
	Debug("Result array", inspectValues(res))
	if len(res) == 0 {
		return nil
	} else if len(res) == 1 {
		// TODO: check CanInterface
		return res[0].Interface()
	} else {
		c := NewCell(nil)
		p := c
		for _, val := range res {
			p = Cons(val.Interface(), p)
		}
		return c
	}
}

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
		exp := Eval(c.Car, s)
		if v := rfl.ValueOf(exp); v.Kind() == rfl.Func {
			if fn, ok := exp.(func(*Cell, *Scope) Any); ok {
				Debug("Found special form function", fn, "in cell", c, "and scope", s)
				return fn(c, s)
			} else {
				fmt.Println(v.Type())
				Debug("Found function", v, "in cell", c, "and scope", s)
				return callGoFunction(v, c, s)
			}
		} else {
			Debug("Expected a function instead of", c.Car)
			return Error(c.Car, " is not a function")
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
