package goutils

func LeftOrRight[T any](condition bool, left, right T) T {
	if condition {
		return left
	}
	return right
}
