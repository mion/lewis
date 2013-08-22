package lewis

import (
	"fmt"
)

func dbg(a interface{}) {
	fmt.Println("-->", a)
}

func Eval(x Any, s *Scope) Any {
	if sym, ok := ToSymbol(x); ok {
		return s.Find(sym).Get(sym) // variable reference
	} else if !IsCell(x) {
		return x // constant literal
	} else {
		c, _ := ToCell(x)
		dbg(c)
		switch c.Cadr(0) {
		case QuoteSymbol:
			return c.Cadr(1)
		default:
			dbg("proc")
			return nil
		}
	}
}
