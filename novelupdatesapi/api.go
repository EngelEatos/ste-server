package novelupdatesapi

import (
	"errors"
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

// Chapter - struct
type Chapter struct {
	ReleaseDate time.Time
	Group       string
	Chapter     string
}

func getBool(s string) (bool, error) {
	s = strings.ToLower(s)
	if s == "yes" {
		return true, nil
	} else if s == "no" {
		return false, nil
	} else {
		return false, fmt.Errorf("wrong string %s - no bool", s)
	}
}

func parseStatus(status string) (int, int, error) {
	exp := regexp.MustCompile(`(?P<chapterCount>\d{1,5})(?:|\+)\sChapters\s\((?P<status>.+?)\)`)
	match := exp.FindStringSubmatch(status)
	chapterCount, err := parseInt(match[1])
	if err != nil {
		return -1, -1, err
	}
	stat, err := getInt(strings.ToLower(match[2]), StatusTypes)
	if err != nil {
		return -1, -1, err
	}
	return chapterCount, stat, nil
}

func strip(s string) string {
	reg := regexp.MustCompile(`\r?\n`)
	s = reg.ReplaceAllString(s, "")
	return strings.Join(strings.Fields(s), "")
}

func parseInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func getSource(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
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
		recommendation[rechtml.Text()] = rechtml.AttrOr("href", "")
	})
	return recommendation
}

func getType(doc *goquery.Document) (int, error) {
	ntype := strings.ReplaceAll(strings.ToLower(doc.Find("a.genre.type").First().Text()), " ", "-")
	rvalue, err := getInt(ntype, NovelTypes)
	if err != nil {
		return -1, err
	}
	return rvalue, nil
}

func getInt(key string, dict map[string]int) (int, error) {
	if val, ok := dict[key]; ok {
		return val, nil
	}
	return -1, fmt.Errorf("%s not in %v", key, dict)
}

func getIntArray(array []string, dict map[string]int) []int {
	var result []int
	for _, key := range array {
		cleanKey := strings.ReplaceAll(strings.ToLower(key), " ", "_")
		if val, ok := dict[cleanKey]; ok {
			result = append(result, val)
		} else {
			log.Printf("%s not in %v\n", key, dict)
		}
	}
	return result
}

func getLanguage(doc *goquery.Document) (int, error) {
	href := doc.Find("div#showlang > a.genre.lang").First().AttrOr("href", "")
	if href == "" {
		return -1, errors.New("href of a.grenre.lang empty")
	}
	lang := href[len("https://www.novelupdates.com/language/") : len(href)-1]
	langid, err := getInt(lang, Languages)
	if err != nil {
		return -1, err
	}
	return langid, nil
}

func getCover(doc *goquery.Document) string {
	coverURL := doc.Find("div.seriesimg > img").First().AttrOr("src", "")
	return coverURL

}

// ParseNovel - parse site , arg novelID
func ParseNovel(novelID string) (*Novel, error) {
	url := baseurl + novelID
	doc, err := getSource(url)
	if err != nil {
		return nil, err
	}
	title := doc.Find(".seriestitlenu").First().Text()
	ntype, err := getType(doc)
	if err != nil {
		log.Printf("failed to get type of novel %s: %s\n", title, err)
	}
	genres := getIntArray(getElementsSlice("div#seriesgenre > a.genre", doc), Genres)
	tags := getIntArray(getElementsSlice("div#showtags > a.genre", doc), Tags)
	lang, err := getLanguage(doc)
	if err != nil {
		log.Printf("failed get langId for novel %s: %s\n", title, err)
	}
	authors := getElementsSlice("div#showauthors > a.genre", doc)
	year, err := parseInt(strip(doc.Find("div#edityear").Text()))
	if err != nil {
		log.Printf("failed to get year of novel %s: %s\n", title, err)
	}
	chapterCount, status, err := parseStatus(doc.Find("div#editstatus").Text())
	if err != nil {
		log.Printf("failed to parse status for %s: %s\n", title, err)
	}
	licensed, err := getBool(strip(doc.Find("div#showlicensed").Text()))
	if err != nil {
		log.Printf("failed to parse bool for %s: %s", title, err)
	}
	completelyTranslated, err := getBool(strip(doc.Find("div#showtranslated").Text()))
	if err != nil {
		log.Printf("failed to parse bool for %s: %s", title, err)
	}
	description := doc.Find("div#editdescription").Text()
	recommendations := getRecommendations(doc)
	coverURL := getCover(doc)

	return &Novel{
		Title:               title,
		Chaptercount:        chapterCount,
		NovelIDSTR:          novelID,
		Type:                ntype,
		Description:         description,
		LanguageID:          lang,
		Year:                year,
		Status:              status,
		Licensed:            licensed,
		CompletlyTranslated: completelyTranslated,
		UpdatedAt:           time.Now(),
		Recommendations:     recommendations,
		Genres:              genres,
		Tags:                tags,
		Authors:             authors,
		CoverURL:            coverURL,
	}, nil
}

// TODO: GetChapter

// GetChapter by novelID
func GetChapter(novelID string, idx int) ([]Chapter, error) {
	doc, err := getSource(baseurl + novelID + "/pg?=" + string(idx))
	if err != nil {
		return nil, err
	}
	var result []Chapter
	doc.Find("table#myTable > tbody > tr").Each(func(idx int, row *goquery.Selection) {
		cols := row.Find("td")
		date := cols.First()
		t, err := time.Parse("01-02-06", date.Text())
		if err != nil {
			log.Println("failed to parse date")
		}
		t = time.Now()

		group := date.Next().AttrOr("href", "")
		chapter := date.Next().Next().AttrOr("href", "")
		result = append(result, Chapter{t, group, chapter})
	})
	return result, nil
}

func getNextPage(doc *goquery.Document) (string, error) {
	np := doc.Find("a.next_page")
	if np.Length() > 0 {
		return np.First().AttrOr("href", ""), nil
	}
	return "", errors.New("next page not found")
}

// TODO: GetLatestSeries()
// func GetLatestSeries() {
// 	idx := 1
// 	url := "https://www.novelupdates.com/latest-series/?st=1&pg=" + string(idx)

// }

// GetNovelsByPage - by idx
func GetNovelsByPage(idx int) (*goquery.Document, bool, error) {
	result := make(map[string]string)
	url := fmt.Sprintf("https://www.novelupdates.com/novelslisting/?st=%d", idx)
	doc, err := getSource(url)
	if err != nil {
		return nil, false, err
	}
	doc.Find("table#myTable > tbody > tr").Each(func(idx int, row *goquery.Selection) {
		col := row.Find("td:nth-child(3) > a")
		result[col.Text()] = col.AttrOr("href", "")
	})
	_, err = getNextPage(doc)
	return doc, err != nil, nil
}

// LiveSearch - post searchtearm to novelupdatesapi livesearch
func LiveSearch(searchTerm string) (string, error) {
	response, err := http.PostForm("https://www.novelupdates.com/wp-admin/admin-ajax.php", url.Values{
		"action":  {"nd_ajaxsearchmain"},
		"strType": {"desktop"},
		"strOne":  {searchTerm},
	})

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
