package lewis

type Symbol struct {
	string
}

var symbols = make(map[string]*Symbol)

func NewSymbol(name string) *Symbol {
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

var QuoteSymbol = NewSymbol("quote")
var IfSymbol = NewSymbol("if")
var SetSymbol = NewSymbol("set!")
var DefineSymbol = NewSymbol("define")
var LambdaSymbol = NewSymbol("lambda")
var BeginSymbol = NewSymbol("begin")

type Scope struct {
	table  map[*Symbol]Any
	parent *Scope
}

func NewScope(parent *Scope) *Scope {
	Scope := new(Scope)
	Scope.table = make(map[*Symbol]Any)
	Scope.parent = parent
	return Scope
}

func (s *Scope) Find(name *Symbol) *Scope {
	for ; s != nil; s = s.parent {
		if _, ok := s.table[name]; ok {
			return s
		}
	}
	return nil
}

func (s *Scope) Set(name *Symbol, val Any) {
	for sc := s; s != nil; s = s.parent {
		if _, ok := sc.table[name]; ok {
			sc.table[name] = val
			return
		}
	}
	s.table[name] = val
}

func (s *Scope) Get(name *Symbol) Any {
	if val, ok := s.table[name]; ok {
		return val
	}
	return nil
}
