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
	var bookTitleList []string
	var bookKeyword string

	fmt.Print("input keyword:")
	fmt.Scan(&search)

	for _, v := range source {
		match, cond := modules.MatchString(search, strings.Split(v, ""))
		if cond {
			bookTitleList = append(bookTitleList, match)
		}
	}

	for i := 0; i < len(bookTitleList); i++ {
		bookKeyword += bookTitleList[i] + " "
	}

	modules.Match(search, bookKeyword, bookTitleList)
}
