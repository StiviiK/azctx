package utils

import "strings"

// StringSlice is a slice of strings, helper type used for extension methods
type StringSlice []string

// LongestStringLength returns the length of the longest string in the given slice
func (slice StringSlice) LongestStringLength() int {
	longestLength := 0
	for _, s := range slice {
		if len(s) > longestLength {
			longestLength = len(s)
		}
	}
	return longestLength
}

// ToLower returns a copy of the given slice of strings with all strings lowercased
func (slice StringSlice) ToLower() StringSlice {
	lowercaseStrings := make(StringSlice, len(slice))
	for i, s := range slice {
		lowercaseStrings[i] = strings.ToLower(s)
	}
	return lowercaseStrings
}
