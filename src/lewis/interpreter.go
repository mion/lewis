package lewis

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
			return fn(c, s)
		} else {
			Debug("Expected a function instead of", c.Car)
			return Error(c.Car, "is not a function")
		}
		// switch c.Cadr(0) {
		// case QuoteSymbol:
		// 	return c.Cadr(1)
		// case IfSymbol:
		// 	Debug("if", c.Cadr(1))
		// 	cond := Eval(c.Cadr(1), s)
		// 	if b, _ := ToBool(cond); b { // TODO: check error
		// 		return Eval(c.Cadr(2), s)
		// 	} else {
		// 		return Eval(c.Cadr(3), s)
		// 	}
		// case SetSymbol:
		// 	variable, _ := ToSymbol(c.Cadr(1)) // TODO: check error
		// 	exp := c.Cadr(2)
		// 	s.Find(variable).Set(variable, Eval(exp, s))
		// 	return nil
		// case BeginSymbol:
		// 	var val Any
		// 	p, _ := c.Cadr(1).(*Cell) // TODO
		// 	Debug("special form", BeginSymbol, "called on", c.Cadr(1))
		// 	for ; !p.IsNull(); p = p.Cdr {
		// 		val = Eval(p.Car, s) // --> Parei aqui: bug quando tenta Eval("if")
		// 	}
		// 	return val
		// default:
		// 	return nil
		// }
	} else {
		return x // literal
	}
}
