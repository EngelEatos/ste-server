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
	ntype := doc.Find("a.genre.type").First().Text()
	var genres []string
	doc.Find("div#seriesgenre").First().Find("a.genre").Each(func(index int, genrehtml *goquery.Selection) {
		genres = append(genres, genrehtml.Text())
	})
	var tags []string
	doc.Find("div#showtags").Find("a.genre").Each(func(index int, taghtml *goquery.Selection) {
		tags = append(tags, taghtml.Text())
	})
	lang := doc.Find("div#showlang > a.genre.lang").Text()
	var authors []string
	doc.Find("div#showauthors > a.genre").Each(func(index int, authorhtml *goquery.Selection) {
		authors = append(authors, authorhtml.Text())
	})
	year := doc.Find("div#edityear").Text()
	status := doc.Find("div#editstatus").Text()
	completelyTranslated := doc.Find("div#showtranslated").Text()
	description := doc.Find("div#editdescription").Text()
	var recommendation map[string]string
	doc.Find("h5.seriesother").First().SiblingsFiltered("a.genre").Each(func(index int, rechtml *goquery.Selection) {
		href, ok := rechtml.Attr("href")
		if !ok {
			href = ""
		}
		recommendation[rechtml.Text()] = href
	})
	fmt.Println(title)
	fmt.Println(ntype)
	fmt.Println(genres)
	fmt.Println(len(genres), genres[0])
	fmt.Println(len(tags), tags[0])
	fmt.Println(lang)
	fmt.Println(authors)
	fmt.Println(year)
	fmt.Println(status)
	fmt.Println(completelyTranslated)
	fmt.Println(description[100])
	fmt.Println(len(recommendation))
}

// GetChapter by novelID
func GetChapter(novelID string) string {
	return ""
}
