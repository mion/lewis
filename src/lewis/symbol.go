package lewis

type Symbol struct {
	string
}

var symbols = make(map[string]*Symbol)

func Sym(name string) *Symbol {
	sym, ok := symbols[name]
	if !ok {
		sym = &Symbol{name}
		symbols[name] = sym
	}
	return sym
}

func (s *Symbol) String() string {
	return s.string
}
