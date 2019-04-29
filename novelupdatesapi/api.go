package novelupdatesapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseurl = "https://www.novelupdates.com/series/"

// ParseNovel - parse site , arg novelID
func ParseNovel(novelID string) {
	url := baseurl + novelID
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find(".seriestitlenu").First().Text()
	fmt.Println(title)
}

// GetChapter by novelID
func GetChapter(novelID string) string {
	return ""
}
