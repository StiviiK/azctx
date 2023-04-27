package utils

import "golang.org/x/exp/constraints"

// Min returns the smaller of x or y.
func Min[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func Max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}
