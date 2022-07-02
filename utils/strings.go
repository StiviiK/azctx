package utils

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
