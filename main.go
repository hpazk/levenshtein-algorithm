package main

import (
	"fmt"

	"github.com/hpazk/levenshtein-algorithm/app/modules"
)

func main() {
	strList := []string{"test", "tester", "tests", "testers", "testing", "tsting", "sting"}
	res, err := modules.FuzzySearch("testign", strList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %s\n", res)
	}
}
