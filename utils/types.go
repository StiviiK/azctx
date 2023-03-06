package utils

// Comparable is an interface that can be used to compare two objects
type Comparable interface {
	Compare(Comparable) int
}

// Named is an interface that can be used to get the name of an object
type Named interface {
	Named() string
}

// ComparableNamed is an interface that can be used to sort a slice of named objects
type ComparableNamed interface {
	Comparable
	Named
}
