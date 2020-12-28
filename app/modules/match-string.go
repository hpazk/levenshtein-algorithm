package modules

import (
	"fmt"
	"strings"
)

// Match is...
func Match(search, keyword string, tmpStore []string) {
	strList := RemoveDuplicateValues(strings.Fields(keyword))

	res, err := FuzzySearch(search, strList)
	if err != nil {
		fmt.Println(err)
	} else if res == "" {
		fmt.Printf("Book not found, search instead for \"%s\"\n", search)
		fmt.Println(tmpStore)
	} else if CompareWord(search, strList) {
		fmt.Println(tmpStore)
	} else {
		fmt.Printf("Did you mean: \"%s\" ?\n", res)
		fmt.Println(tmpStore)
	}
	fmt.Println(len(tmpStore))
}
