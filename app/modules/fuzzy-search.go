package modules

// StringsSimilarity is...
func stringsSimilarity(str1 string, str2 string) (float32, error) {
	return matchingIndex(str1, str2, levenshteinDistance(str1, str2)), nil
}

// Return matching index E [0..1] from two strings and an edit distance
func matchingIndex(str1 string, str2 string, distance int) float32 {
	// Compare strings length and make a matching percentage between them
	if len(str1) >= len(str2) {
		return float32(len(str1)-distance) / float32(len(str1))
	}
	return float32(len(str2)-distance) / float32(len(str2))
}

// FuzzySearch realize an approximate search on a string list and return the closest one compared
// to the string input
func FuzzySearch(str string, strList []string) (string, error) {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range strList {
		sim, err := stringsSimilarity(str, strToCmp)
		if err != nil {
			return "", err
		}

		if sim == 1.0 {
			return strToCmp, nil
		} else if sim > higherMatchPercent {
			higherMatchPercent = sim
			tmpStr = strToCmp
		}
	}

	return tmpStr, nil
}
