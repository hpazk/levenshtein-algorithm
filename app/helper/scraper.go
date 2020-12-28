package helper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// BookScraper is...
func Scraper() []string {
	url := "https://www.penguin.co.uk/articles/2018/100-must-read-classic-books.html"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("error: %d, status:, %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var title []string

	// var title []string
	doc.Find(".bookCard__wrapper").Each(func(i int, sel *goquery.Selection) {
		text := sel.Find("h5").Text()
		title = append(title, text)
	})

	return title
}
