package lewis

var SpecialFormsLibrary = &Library{functions: map[string]Any{
	"import": func(c *Cell, s *Scope) Any {
		name, ok := c.Cadr(1).(*Symbol)
		if !ok {
			Error("import -- argument is not a Symbol")
		}
		if lib, ok := Libraries[name.String()]; ok {
			s.Import(lib)
		} else {
			Error("import -- no library named ", name)
		}
		return nil
	},
	"int64": func(c *Cell, s *Scope) Any {
		return ToInt64(c.Cadr(1))
	},
	"quote": func(c *Cell, s *Scope) Any {
		return c.Cadr(1)
	},
	"if": func(c *Cell, s *Scope) Any {
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
	"set!": func(c *Cell, s *Scope) Any {
		variable, _ := c.Cadr(1).(*Symbol) // TODO: check error
		exp := c.Cadr(2)
		Debug("Setting variable", variable, "to expression", exp)
		s.Find(variable).Set(variable, Eval(exp, s))
		return nil
	},
	"define": func(c *Cell, s *Scope) Any {
		Debug("define form called on cell", c, "in scope", s)
		variable, _ := c.Cadr(1).(*Symbol) // TODO: check error
		exp := Eval(c.Cadr(2), s)
		Debug("Define", variable, "as expression", exp)
		s.Set(variable, exp)
		return nil
	},
	"lambda": func(c *Cell, s *Scope) Any {
		cellArgs := c.Cadr(1).(*Cell) // TODO: check
		args := make([]*Symbol, cellArgs.Len())
		for i := 0; i < cellArgs.Len(); i++ {
			args[i] = cellArgs.Cadr(i).(*Symbol) // TODO: check type
		}
		body := c.Cadr(2) // TODO: Copy?
		Debug("lambda form with args", args, "and body", body)
		return func(cell *Cell, scope *Scope) Any {
			local := NewScope(scope) // TODO: check possibly erroneous lexical scoping
			for i, arg := range args {
				local.Set(arg, cell.Cadr(i+1))
			}
			return Eval(body, local)
		}
	},
	"begin": func(c *Cell, s *Scope) Any {
		var res Any
		for i := 1; i < c.Len(); i++ {
			res = Eval(c.Cadr(i), s)
		}
		return res
	},
	// Temporary
	// "+": func(c *Cell, s *Scope) Any {
	// 	Debug("adding ", c.Cadr(1), "to", c.Cadr(2))
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	Debug("converted to int", a, "and ", b, "resulted", a+b)
	// 	return a + b
	// },
	// "-": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a - b
	// },
	// "<": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a < b
	// },
	// ">": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a > b
	// },
	// "=": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a == b
	// },
	// "*": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a * b
	// },
	// "/": func(c *Cell, s *Scope) Any {
	// 	a, _ := Eval(c.Cadr(1), s).(int)
	// 	b, _ := Eval(c.Cadr(2), s).(int)
	// 	return a / b
	// },
}}
