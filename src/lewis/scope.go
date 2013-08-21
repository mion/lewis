package lewis

type Scope struct {
	table  map[string]Any
	parent *Scope
}

func NewScope(parent *Scope) *Scope {
	Scope := new(Scope)
	Scope.table = make(map[string]Any)
	Scope.parent = parent
	return Scope
}

func (s *Scope) Find(name string) *Scope {
	for ; s != nil; s = s.parent {
		if _, ok := s.table[name]; ok {
			return s
		}
	}
	return nil
}

func (s *Scope) Set(name string, val Any) {
	for sc := s; s != nil; s = s.parent {
		if _, ok := sc.table[name]; ok {
			sc.table[name] = val
			return
		}
	}
	s.table[name] = val
}

func (s *Scope) Get(name string) Any {
	if val, ok := s.table[name]; ok {
		return val
	}
	return nil
}
