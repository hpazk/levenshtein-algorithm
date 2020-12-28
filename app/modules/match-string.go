package modules

import (
	"fmt"
	"strings"
)

// Match is...
func Match(search, keyword string) {
	strList := RemoveDuplicateValues(strings.Fields(keyword))

	res, err := FuzzySearch(search, strList)
	if err != nil {
		fmt.Println(err)
	} else if res == "" {
		fmt.Printf("Search instead for \"%s\"\n", search)
		fmt.Println(strList)
	} else if CompareWord(search, strList) {
		fmt.Println(strList)
	} else {
		fmt.Printf("Did you mean: \"%s\" ?\n", res)
		fmt.Println(strList)
	}
}
