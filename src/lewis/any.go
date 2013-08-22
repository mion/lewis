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

func ToBool(a Any) (bool, bool) {
	if b, ok := a.(bool); ok {
		return b, true
	} else {
		return false, false
	}
}

func AsString(a Any) (string, bool) {
	if s, ok := a.(string); ok {
		return s, true
	} else {
		return "", false
	}
}
