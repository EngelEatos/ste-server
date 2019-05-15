package novelupdatesapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// TODO: Errorhandling, dont die if smth fail

var baseurl = "https://www.novelupdates.com/series/"

// Novel - struct
type Novel struct {
	Title               string
	Chaptercount        int
	NovelIDSTR          string
	Type                int
	Genres              []int
	Tags                []int
	Description         string
	LanguageID          int
	Year                int
	Status              int
	Licensed            bool
	CompletlyTranslated bool
	UpdatedAt           time.Time
	Recommendations     map[string]string
	Authors             []string
	CoverURL            string
}

func getBool(s string) bool {
	s = strings.ToLower(s)
	if s == "yes" {
		return true
	} else if s == "no" {
		return false
	} else {
		log.Fatalf("wrong string %s - no bool\n", s)
		return false
	}
}

func parseStatus(status string) map[string]int {
	exp := regexp.MustCompile(`(?P<chapterCount>\d{1,5})(?:|\+)\sChapters\s\((?P<status>.+?)\)`)
	match := exp.FindStringSubmatch(status)
	return map[string]int{
		"chapterCount": parseInt(match[1]),
		"status":       getInt(strings.ToLower(match[2]), StatusTypes),
	}
}

func strip(s string) string {
	reg := regexp.MustCompile(`\r?\n`)
	s = reg.ReplaceAllString(s, "")
	return strings.Join(strings.Fields(s), "")
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func getSource(url string) *goquery.Document {
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
	return doc
}

func getElementsSlice(selector string, doc *goquery.Document) []string {
	var result []string
	doc.Find(selector).Each(func(index int, html *goquery.Selection) {
		result = append(result, html.Text())
	})
	return result
}

func getRecommendations(doc *goquery.Document) map[string]string {
	recommendation := make(map[string]string)
	doc.Find("div.wpb_wrapper > a.genre").Each(func(index int, rechtml *goquery.Selection) {
		href, ok := rechtml.Attr("href")
		if !ok {
			href = ""
		}
		recommendation[rechtml.Text()] = href
	})
	return recommendation
}

func getType(doc *goquery.Document) int {
	ntype := strings.ReplaceAll(strings.ToLower(doc.Find("a.genre.type").First().Text()), " ", "-")
	return getInt(ntype, NovelTypes)
}

func getInt(key string, dict map[string]int) int {
	if val, ok := dict[key]; ok {
		return val
	}
	log.Fatalf("%s not in %v\n", key, dict)
	return -1
}

func getIntArray(array []string, dict map[string]int) []int {
	var result []int
	for _, key := range array {
		cleanKey := strings.ReplaceAll(strings.ToLower(key), " ", "_")
		if val, ok := dict[cleanKey]; ok {
			result = append(result, val)
		} else {
			log.Fatalf("%s not in %v\n", key, dict)
		}
	}
	return result
}

func getLanguage(doc *goquery.Document) int {
	href, ok := doc.Find("div#showlang > a.genre.lang").First().Attr("href")
	if !ok {
		log.Fatal("no href attribute on a.genre.lang")
	}
	lang := href[len("https://www.novelupdates.com/language/") : len(href)-1]
	return getInt(lang, Languages)
}

func getCover(doc *goquery.Document) string {
	coverURL, ok := doc.Find("div.seriesimg > img").First().Attr("src")
	if !ok {
		log.Fatal("no src attribute on div.seriesimg > img")
	}
	return coverURL

}

// TODO: Errorhandling

// ParseNovel - parse site , arg novelID
func ParseNovel(novelID string) Novel {
	url := baseurl + novelID
	doc := getSource(url)
	title := doc.Find(".seriestitlenu").First().Text()
	ntype := getType(doc)
	genres := getIntArray(getElementsSlice("div#seriesgenre > a.genre", doc), Genres)
	tags := getIntArray(getElementsSlice("div#showtags > a.genre", doc), Tags)
	lang := getLanguage(doc)
	authors := getElementsSlice("div#showauthors > a.genre", doc)
	year := parseInt(strip(doc.Find("div#edityear").Text()))
	status := parseStatus(doc.Find("div#editstatus").Text())
	licensed := getBool(strip(doc.Find("div#showlicensed").Text()))
	completelyTranslated := getBool(strip(doc.Find("div#showtranslated").Text()))
	description := doc.Find("div#editdescription").Text()
	recommendations := getRecommendations(doc)
	coverURL := getCover(doc)

	return Novel{
		Title:               title,
		Chaptercount:        status["chapterCount"],
		NovelIDSTR:          novelID,
		Type:                ntype,
		Description:         description,
		LanguageID:          lang,
		Year:                year,
		Status:              status["status"],
		Licensed:            licensed,
		CompletlyTranslated: completelyTranslated,
		UpdatedAt:           time.Now(),
		Recommendations:     recommendations,
		Genres:              genres,
		Tags:                tags,
		Authors:             authors,
		CoverURL:            coverURL,
	}
}

// TODO: GetChapter

// GetChapter by novelID
func GetChapter(novelID string) string {
	return ""
}

// TODO: finish
func getPage(idx int) (*goquery.Document, int, error) {
	url := fmt.Sprintf("https://www.novelupdates.com/novelslisting/?st=%d", idx)
	source := getSource(url)
	return source, 0, nil
}

// TODO: Errorhandling

// LiveSearch ...
func LiveSearch(searchTerm string) string {
	response, err := http.PostForm("https://www.novelupdates.com/wp-admin/admin-ajax.php", url.Values{
		"action":  {"nd_ajaxsearchmain"},
		"strType": {"desktop"},
		"strOne":  {searchTerm},
	})

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		//handle read response error
		log.Fatal(err)
	}
	return string(body)
}
