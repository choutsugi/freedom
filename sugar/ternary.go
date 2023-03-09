package sugar

func Ternary[T any](condition bool, ifRes, elseRes T) T {
	if condition {
		return ifRes
	}
	return elseRes
}
