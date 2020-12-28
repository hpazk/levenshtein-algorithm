package helper

// Compare is...
func Compare(a, b []rune) bool {
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

// Min is...
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
