package lewis

import (
	rfl "reflect"
	"fmt"
)

type Any interface{}

func ToInt64(a Any) int64 {
	switch v := a.(type) {
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	default:
		return int64(0)
	}
}

func Inspect(a Any) {
	fmt.Println("Type:", rfl.TypeOf(a))
}
