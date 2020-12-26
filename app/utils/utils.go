package utils

// EqualCompare is...
func EqualCompare(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// FindMin is...
func FindMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
