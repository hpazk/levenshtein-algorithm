package modules

import (
	"fmt"

	"github.com/hpazk/levenshtein-algorithm/app/utils"
)

func levenshteinDistance(t, s string) int {
	target := []rune(t)
	source := []rune(s)

	fmt.Println("rune target", target)
	fmt.Println("rune source", source)

	targetLen := len(target)
	sourceLen := len(source)

	fmt.Println("len target", targetLen)
	fmt.Println("len source", sourceLen)

	if targetLen == 0 {
		return sourceLen
	} else if sourceLen == 0 {
		return targetLen
	} else if utils.EqualCompare(target, source) {
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
			col[i] = utils.FindMin(
				utils.FindMin(col[i]+1, col[i-1]+1),
				lastKey+x,
			)
			lastKey = oldKey
		}
	}

	fmt.Println(col[targetLen])

	return col[targetLen]
}
