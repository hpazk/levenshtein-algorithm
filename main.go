package main

import (
	"fmt"
	"strings"

	"github.com/hpazk/levenshtein-algorithm/app/helper"
	"github.com/hpazk/levenshtein-algorithm/app/modules"
)

func main() {
	source := helper.Scraper()
	var search string

	fmt.Print("input keyword:")
	fmt.Scan(&search)

	var bookTitle []string

	for _, v := range source {
		match, cond := modules.MatchString(search, strings.Split(v, ""))
		if cond {
			bookTitle = append(bookTitle, match)
		}
	}

	var bookKeyword string
	for i := 0; i < len(bookTitle); i++ {
		// fmt.Println(result[i])

		bookKeyword += bookTitle[i] + " "
	}

	modules.Match(search, bookKeyword)

	// strList0 := []string{"test", "trigger", "top", "ten", "throw", "tonight"}
	// for i := 0; i < len(result); i++ {

	// }
}
