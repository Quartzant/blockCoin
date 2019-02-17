package apiService

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (as *apiService) Letter20() (r []*LettersResultOutput, err error) {
	var result []*LettersResultOutput
	res, err := http.Get(as.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#app > div.main.bbt-main--margin.bbt-main > div > div > div.bbt-col-xs-16 > div.flash-content > div > div.flash-module > div > div.flash-module__lists > ul > li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		result = append(result, &LettersResultOutput{
			Time: s.Find("div.flash-item__body > span").Text(),
			Title: s.Find("div.flash-item__body > a").Text(),
			Content: s.Find("div.flash-item__body > div").Text(),
			Up: s.Find("div.operate-box.bbt-clearfix > span.operate-item.operate-item__up").Text(),
			Down: s.Find("div.operate-box.bbt-clearfix > span.operate-item.operate-item__down").Text(),
		})
	})
	return result, nil
}
