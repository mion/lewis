package lewis

import (
	"fmt"
)

var Debugging = true

// func dbg(format string, a ...interface{}) {
// 	if Debugging {
// 		fmt.Printf("[ DEBUG ]\n\t"+format+"\n\n", a...)
// 	}
// }

func debug(a ...interface{}) {
	fmt.Println("Debugger:")
	for i, v := range a {
		if i%2 == 0 {
			fmt.Print("\t")
		} else {
			fmt.Print("\t\t")
		}
		fmt.Println(v)
	}
}

func Error(msg string) {
	fmt.Println("[!] Error:", msg)
}

func Eval(x Any, s *Scope) Any {
	// dbg("evaluating \n\t\t%v\n\tin scope\n\t\t%v", x, s)
	debug("evaluating", x, "in scope", s)
	if sym, ok := ToSymbol(x); ok {
		debug("found symbol", x)
		if scope := s.Find(sym); scope != nil {
			return scope.Get(sym) // variable reference
		} else {
			Error("undefined symbol")
			return nil
		}
	} else if !IsCell(x) {
		return x // constant literal
	} else {
		c, _ := ToCell(x) // TODO: check error
		switch c.Cadr(0) {
		case QuoteSymbol:
			return c.Cadr(1)
		case IfSymbol:
			debug("if", c.Cadr(1))
			cond := Eval(c.Cadr(1), s)
			if b, _ := ToBool(cond); b { // TODO: check error
				return Eval(c.Cadr(2), s)
			} else {
				return Eval(c.Cadr(3), s)
			}
		case SetSymbol:
			variable, _ := ToSymbol(c.Cadr(1)) // TODO: check error
			exp := c.Cadr(2)
			s.Find(variable).Set(variable, Eval(exp, s))
			return nil
		case BeginSymbol:
			var val Any
			p, _ := c.Cadr(1).(*Cell) // TODO
			debug("begin", c.Cadr(1), "with", p.Car)
			for ; !p.IsNull(); p = p.Cdr {
				debug("begin: loop", p.Car)
				val = Eval(p.Car, s) // --> Parei aqui: bug quando tenta Eval("if")
			}
			return val
		default:
			return nil
		}
	}
}
