package lewis

import (
	"fmt"
)

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

type Scope struct {
	table  map[*Symbol]Any
	parent *Scope
}

var GlobalScope = NewScope(nil).Import(SpecialFormsLibrary)

func NewScope(parent *Scope) *Scope {
	Scope := new(Scope)
	Scope.table = make(map[*Symbol]Any)
	Scope.parent = parent
	return Scope
}

func (s *Scope) String() string {
	// if s.parent == nil {
	// 	return "Global"
	// }
	str := "{ "
	for key, val := range s.table {
		str += fmt.Sprintf("%v: %v, ", key, val)
	}
	str += " } --> " + fmt.Sprint(s.parent)
	return str
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
	// for sc := s; s != nil; sc = sc.parent {
	// 	if _, ok := sc.table[name]; ok {
	// 		sc.table[name] = val
	// 		return
	// 	}
	// }
	s.table[name] = val
}

func (s *Scope) Get(name *Symbol) Any {
	if val, ok := s.table[name]; ok {
		return val
	}
	return nil
}

func (s *Scope) Import(l *Library) *Scope {
	// TODO: check for conflicts
	Debug("Importing library")
	for name, fn := range l.functions {
		s.table[Sym(name)] = fn
	}
	return s
}
