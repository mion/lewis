package lewis

var OperatorsLibrary = &Library{functions: map[string]Any{
	"+": func(c *Cell, s *Scope) Any {
		result := int64(0)
		for i := 1; i < c.Len(); i++ {
			result += Eval(c.Cadr(i), s).(int64)
		}
		return result
	},
	"-": func(c *Cell, s *Scope) Any {
		result := int64(0)
		for i := 1; i < c.Len(); i++ {
			result -= Eval(c.Cadr(i), s).(int64)
		}
		return result
	},
	"*": func(c *Cell, s *Scope) Any {
		result := int64(1)
		for i := 1; i < c.Len(); i++ {
			result *= Eval(c.Cadr(i), s).(int64)
		}
		return result
	},
	"/": func(c *Cell, s *Scope) Any {
		result := int64(1)
		for i := 1; i < c.Len(); i++ {
			result /= Eval(c.Cadr(i), s).(int64)
		}
		return result
	},
	"=": func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int64)
		b, _ := Eval(c.Cadr(2), s).(int64)
		return a == b
	},
	"<": func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int64)
		b, _ := Eval(c.Cadr(2), s).(int64)
		return a < b
	},
	">": func(c *Cell, s *Scope) Any {
		a, _ := Eval(c.Cadr(1), s).(int64)
		b, _ := Eval(c.Cadr(2), s).(int64)
		return a > b
	},
}}
