package lewis

type Any interface{}

func ToSymbol(a Any) (*Symbol, bool) {
	if sym, ok := a.(*Symbol); ok {
		return sym, true
	} else {
		return nil, false
	}
}

func IsSymbol(a Any) bool {
	_, ok := ToSymbol(a)
	return ok
}

func ToCell(a Any) (*Cell, bool) {
	if cell, ok := a.(*Cell); ok {
		return cell, true
	} else {
		return nil, false
	}
}

func IsCell(a Any) bool {
	_, ok := ToCell(a)
	return ok
}
