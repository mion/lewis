package lewis

var GlobalScope = &Scope{map[*Symbol]Any{
	Sym("quote"): func(c *Cell, s *Scope) Any {
		return c.Cadr(1)
	},
	Sym("if"): func(c *Cell, s *Scope) Any {
		switch cond := Eval(c.Cadr(1), s).(type) {
		case bool:
			if cond {
				return Eval(c.Cadr(2), s)
			} else {
				return Eval(c.Cadr(3), s)
			}
		default:
			return Error(cond, "used as 'if' condition")
		}
	},
	Sym("set!"): func(c *Cell, s *Scope) Any {
		variable, _ := c.Cadr(1).(*Symbol) // TODO: check error
		exp := c.Cadr(2)
		Debug("Setting variable", variable, "to expression", exp)
		s.Find(variable).Set(variable, Eval(exp, s))
		return nil
	},
	Sym("define"): func(c *Cell, s *Scope) Any {
		Debug("define form called on cell", c, "in scope", s)
		variable, _ := c.Cadr(1).(*Symbol) // TODO: check error
		exp := Eval(c.Cadr(2), s)
		Debug("Define", variable, "as expression", exp)
		s.Set(variable, exp)
		return nil
	},
	// Temporary
	Sym("+"): func(c *Cell, s *Scope) Any {
		Debug("adding ", c.Cadr(1), "to", c.Cadr(2))
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		Debug("converted to int", a, "and ", b, "resulted", a+b)
		return a + b
	},
	Sym("-"): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a - b
	},
	Sym("<"): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a < b
	},
	Sym(">"): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a > b
	},
	Sym("="): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a == b
	},
	Sym("*"): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a * b
	},
	Sym("/"): func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int)
		b, _ := Eval(c.Cadr(2), s).(int)
		return a / b
	},
}, nil}
