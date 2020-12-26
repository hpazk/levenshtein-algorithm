package modules

func matchingIndex(t, s string) (float32, error) {
	if len(t) >= len(s) {
		return float32(len(t) - levenshteinDistance(t, s)/len(t)), nil
	}

	return float32(len(s) - levenshteinDistance(t, s)/len(s)), nil
}

func fuzzySearch(t string, sourceList []string) (string, error) {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range sourceList {
		sim, err := matchingIndex(t, strToCmp)
		if err != nil {
			return "", nil
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
