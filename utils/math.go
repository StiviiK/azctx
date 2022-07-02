package utils

import "golang.org/x/exp/constraints"

// Min returns the smaller of x or y.
func Min[T constraints.Integer](x, y T) T {
	if x > y {
		return y
	}
	return x
}
