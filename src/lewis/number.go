package lewis

func ConvertInt64(a Any) int64 {
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
