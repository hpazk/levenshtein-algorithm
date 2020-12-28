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
	var strList []string

	fmt.Print("Search book title:")
	fmt.Scan(&search)

	var result []string

	for _, v := range source {
		match, cond := modules.MatchString(search, strings.Split(v, ""))
		if cond {
			result = append(result, match)
		}
	}

	// Store to multidimentional array
	tmpStore := make([][]string, len(result))

	for i := 0; i < len(result); i++ {
		tmpStore[i] = append(tmpStore[i], result[i])
	}

	var tn string
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])

		tn += result[i] + " "
	}

	// strList0 := []string{"test", "trigger", "top", "ten", "throw", "tonight"}
	strList = strings.Fields(tn)
	res, err := modules.FuzzySearch(search, strList)
	if err != nil {
		fmt.Println(err)
	} else if res == "" {
		fmt.Printf("Book title \"%s\" not found\n", search)
	} else {
		fmt.Printf("Did you mean: \"%s\" ?\n", res)
	}

	fmt.Println(result)
}
