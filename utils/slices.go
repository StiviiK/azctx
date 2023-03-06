package utils

// NamedSortableSlice is an interface that can be used to sort a slice of named objects
type ComparableNamedSlice[T ComparableNamed] []T

// Implement the sort.Interface interface
func (slice ComparableNamedSlice[T]) Len() int           { return len(slice) }
func (slice ComparableNamedSlice[T]) Less(i, j int) bool { return slice[i].Compare(slice[j]) < 0 }
func (slice ComparableNamedSlice[T]) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

// Names returns the names of the give slice of named objects
func (slice ComparableNamedSlice[T]) Names() (ret []string) {
	for _, subscription := range slice {
		ret = append(ret, subscription.Named())
	}
	return
}

// Filter filters a slice of any type using a test function
func (slice ComparableNamedSlice[T]) Filter(test func(T) bool) (ret ComparableNamedSlice[T]) {
	for _, s := range slice {
		if ok := test(s); ok {
			ret = append(ret, s)
		}
	}
	return
}
