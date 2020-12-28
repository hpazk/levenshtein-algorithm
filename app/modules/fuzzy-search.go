package modules

import (
	"strings"
)

func matchingIndex(t, s string) (float32, error) {
	distance := levenshteinDistance(t, s)
	if len(t) >= len(t) {
		return float32(len(t)-distance) / float32(len(t)), nil
	}
	return float32(len(s)-distance) / float32(len(s)), nil
}

// func stringsSimilarity(t, s string) (float32, error) {
// 	return matchingIndex(t, s, levenshteinDistance(t, s)), nil
// }

// MatchString is...
func MatchString(str string, s []string) (string, bool) {
	trimStr := strings.ToLower(strings.ReplaceAll(str, " ", ""))
	var tokens []string

	var sPos = 0

	for i := 0; i < len(s); i++ {
		chr := strings.ToLower(s[i])

		if sPos < len(str) && chr == string(trimStr[sPos]) {
			tmpChr := chr
			chr = tmpChr
			sPos++
		}
		tokens = append(tokens, chr)
	}

	if sPos != len(str) {
		return "", false
	}

	result := strings.Join(tokens, "")

	return result, true
}

// FuzzySearch is...
func FuzzySearch(t string, sList []string) (string, error) {
	var higherMatchPercent float32
	var tmpStr string
	for _, strToCmp := range sList {
		sim, err := matchingIndex(t, strToCmp)
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
