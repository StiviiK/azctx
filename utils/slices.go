package utils

// Filter filters a slice of any type using a test function
func Filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, s := range slice {
		if ok := test(s); ok {
			ret = append(ret, s)
		}
	}
	return
}
