package utils

import "strings"

// GetLongestStringLength returns the length of the longest string in the given slice
func GetLongestStringLength(strings []string) int {
	longestLength := 0
	for _, s := range strings {
		if len(s) > longestLength {
			longestLength = len(s)
		}
	}
	return longestLength
}

// LowercaseStrings returns a copy of the given slice of strings with all strings lowercased
func LowercaseStrings(input []string) []string {
	lowercaseStrings := make([]string, len(input))
	for i, s := range input {
		lowercaseStrings[i] = strings.ToLower(s)
	}
	return lowercaseStrings
}
