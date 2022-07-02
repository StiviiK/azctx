package utils

// RemoveUTF8BOM removes the UTF8 BOM from the given byte array if present
func RemoveUTF8BOM(data []byte) []byte {
	if len(data) > 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return data[3:]
	}
	return data
}
