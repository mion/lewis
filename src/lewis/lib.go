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
		s.Find(variable).Set(variable, Eval(exp, s))
		return nil
	},
	Sym("define"): func(c *Cell, s *Scope) Any {
		variable, _ := c.Cadr(1).(*Symbol) // TODO: check error
		exp := c.Cadr(2)
		s.Set(variable, Eval(exp, s))
		return nil
	},
	// Temporary
	Sym("+"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a + b
	},
	Sym("-"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a - b
	},
	Sym("<"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a < b
	},
	Sym(">"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a > b
	},
	Sym("="): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a == b
	},
	Sym("*"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a * b
	},
	Sym("/"): func(c *Cell, s *Scope) Any {
		a, _ := c.Cadr(1).(int64)
		b, _ := c.Cadr(2).(int64)
		return a / b
	},
}, nil}
