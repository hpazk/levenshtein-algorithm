package modules

import (
	"fmt"
)

// Levenshtein is...
func Levenshtein(t, s string) int {
	source := []rune(s)
	target := []rune(t)

	fmt.Println("source", source)
	fmt.Println("target", target)
	return 0
}
