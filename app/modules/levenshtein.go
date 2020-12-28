package modules

import (
	"github.com/hpazk/levenshtein-algorithm/app/helper"
)

func levenshteinDistance(t, s string) int {
	target := []rune(t)
	source := []rune(s)

	targetLen := len(target)
	sourceLen := len(source)

	if targetLen == 0 {
		return sourceLen
	} else if sourceLen == 0 {
		return targetLen
	} else if helper.Compare(target, source) {
		return 0
	}

	col := make([]int, targetLen+1)

	for i := 1; i <= targetLen; i++ {
		col[i] = i
	}

	for j := 1; j <= sourceLen; j++ {
		col[0] = j
		lastKey := j - 1
		for i := 1; i <= targetLen; i++ {
			oldKey := col[i]
			var x int
			if target[i-1] != source[j-1] {
				x = 1
			}
			col[i] = helper.Min(
				helper.Min(col[i]+1, col[i-1]+1),
				lastKey+x,
			)
			lastKey = oldKey
		}
	}

	return col[targetLen]
}
