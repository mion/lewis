package lewis

import (
	"fmt"
	"reflect"
)

type Atom struct {
	IsLiteral bool
	Value     Any
}

func (a *Atom) String() string {
	if a.IsLiteral {
		return fmt.Sprintf("{%v: %v}", reflect.TypeOf(a.Value), a.Value)
	} else {
		return fmt.Sprintf("{identifier: %v}", a.Value)
	}
}

func MakeLiteral(v Any) *Atom {
	return &Atom{IsLiteral: true, Value: v}
}

func MakeIdentifier(v Any) *Atom {
	return &Atom{IsLiteral: false, Value: v}
}
