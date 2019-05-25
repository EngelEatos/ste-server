package novelupdatesapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"ste/utils"
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
	FetchedAt           time.Time
	Recommendations     map[string]string
	Authors             []string
	CoverURL            string
}

// TODO: compare function between model.Novel and nuapi.Novel
// func (n *Novel) equal(ctx context.Context, novel *models.Novel, db *sql.DB) bool {
// 	genres, err := novel.Genres().All(ctx, db)
// 	if err != nil {

// 	}
// 	var intGenres []int
// 	for _, genre := range genres {
// 		intGenres = append(intGenres, genre.ID)
// 	}
// 	tags, err := novel.Tags().All(ctx, db)
// 	var intTags []int
// 	for _, tag := range tags {
// 		intTags = append(intTags, tag.ID)
// 	}

// 	authors, err := novel.Authors().All(ctx, db)
// 	if err != nil {

// 	}
// 	var strAuthors []string
// 	for _, author := range authors {
// 		strAuthors = append(strAuthors, author.Name.String)
// 	}
// 	cover, err := novel.Cover().One(ctx, db)
// 	if err != nil {

// 	}
// 	ntype, err := novel.Ntype().One(ctx, db)
// 	if err != nil {

// 	}
// 	return (n.Title == novel.Title) && (n.Chaptercount == novel.Chaptercount.Int) && (n.NovelIDSTR == novel.NovelIDSTR.String) &&
// 		(n.Type == ntype.ID) && (reflect.DeepEqual(novel.Genres, intGenres) && (reflect.DeepEqual(novel.Tags, intTags)) &&
// 		(n.Description == novel.Description.String) && (n.LanguageID == novel.LanguageID.Int) && (n.Year == novel.Year.Int) && (n.Status == novel.Status.Int) &&
// 		(n.Licensed == novel.Licensed.Bool) && (n.CompletlyTranslated == novel.CompletlyTranslated.Bool) && (reflect.DeepEqual(n.Authors, strAuthors)) &&
// 		(n.CoverURL == cover.URL))

// }

// Group struct
type Group struct {
	Name string
	URL  string
}

// Chapter - struct
type Chapter struct {
	ReleaseDate time.Time
	Group       Group
	Title       string
	URL         string
	Idx         int
	EndIdx      int
	Part        int
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
	chapterCount, err := utils.ParseInt(match[1])
	if err != nil {
		return -1, -1, err
	}
	stat, err := utils.GetInt(strings.ToLower(match[2]), StatusTypes)
	if err != nil {
		return -1, -1, err
	}
	return chapterCount, stat, nil
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
	rvalue, err := utils.GetInt(ntype, NovelTypes)
	if err != nil {
		return -1, err
	}
	return rvalue, nil
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
	langid, err := utils.GetInt(lang, Languages)
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
	doc, err := utils.GetSource(url)
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
	year, err := utils.ParseInt(utils.Strip(doc.Find("div#edityear").Text()))
	if err != nil {
		log.Printf("failed to get year of novel %s: %s\n", title, err)
	}
	chapterCount, status, err := parseStatus(doc.Find("div#editstatus").Text())
	if err != nil {
		log.Printf("failed to parse status for %s: %s\n", title, err)
	}
	licensed, err := getBool(utils.Strip(doc.Find("div#showlicensed").Text()))
	if err != nil {
		log.Printf("failed to parse bool for %s: %s", title, err)
	}
	completelyTranslated, err := getBool(utils.Strip(doc.Find("div#showtranslated").Text()))
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
		FetchedAt:           time.Now(),
		Recommendations:     recommendations,
		Genres:              genres,
		Tags:                tags,
		Authors:             authors,
		CoverURL:            coverURL,
	}, nil
}

// ResolveNuLinks -- resolve intern novelupdates links
func ResolveNuLinks(url string) string {
	if !strings.HasPrefix(url, "https://www.novelupdates.com/extnu/") {
		return url
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	return res.Request.URL.String()
}

// GetMaxPage --
func GetMaxPage(novelID string) (int, error) {
	doc, err := utils.GetSource(baseurl + novelID)
	if err != nil {
		return -1, err
	}
	pagination := doc.Find("div.digg_pagination > a:nth-child(5)")
	if pagination.Length() > 0 {
		return utils.ParseInt(pagination.First().Text())
	}
	return -1, errors.New("pagination not found")
}

// GetAllChapter - GetAllChapter by novelID
func GetAllChapter(novelID string) ([]Chapter, error) {
	var result []Chapter
	lastPage, err := GetMaxPage(novelID)
	if err != nil {
		return result, err
	}
	for i := 1; i <= lastPage; i++ {
		chapter, err := GetChapter(novelID, i)
		if err != nil {
			return result, err
		}
		result = append(result, chapter...)
	}
	return result, nil
}

// GetChapter by novelID, idx
func GetChapter(novelID string, idx int) ([]Chapter, error) {
	url := fmt.Sprintf("%s%s/?pg=%d", baseurl, novelID, idx)
	log.Println(url)
	doc, err := utils.GetSource(url)
	if err != nil {
		return nil, err
	}
	var result []Chapter
	doc.Find("table#myTable > tbody > tr").Each(func(idx int, row *goquery.Selection) {
		c := Chapter{}
		row.Find("td").Each(func(idx int, col *goquery.Selection) {
			if idx == 0 {
				// ReleaseDate
				t, err := time.Parse("01/02/06", col.Text())
				if err != nil {
					log.Println("failed to parse date")
				}
				c.ReleaseDate = t
			} else if idx == 1 {
				// group
				gurl := col.Find("a").First().AttrOr("href", "")
				group := Group{Name: strings.TrimSpace(col.Text()), URL: gurl}
				c.Group = group
			} else if idx == 2 {
				// chapter url
				itype, sIdx := parseIdx(col.Text())
				if itype == 0 {
					// chapter range
					c.Idx, err = utils.ParseInt(sIdx[0])
					if err != nil {
						log.Printf("failed to parse Idx: %s\n", sIdx[0])
					}
					c.EndIdx, err = utils.ParseInt(sIdx[1])
					if err != nil {
						log.Printf("failed to parse EndIdx: %s\n", sIdx[1])
					}
					// TODO: resolve range issue
				} else if itype == 1 {
					if len(sIdx) == 1 {
						// c99
						c.Idx, err = utils.ParseInt(sIdx[0])
						if err != nil {
							log.Printf("failed to parse Idx: %s\n", sIdx[0])
						}
						c.Title = "Chapter " + sIdx[0]
					} else if len(sIdx) == 2 {
						// c99.1 - c99 part 1
						c.Idx, err = utils.ParseInt(sIdx[0])
						if err != nil {
							log.Printf("failed to parse Idx: %s\n", sIdx[0])
						}
						c.Part, err = utils.ParseInt(sIdx[1])
						if err != nil {
							log.Printf("failed to parse Part: %s\n", sIdx[1])
						}
						c.Title = "Chapter " + sIdx[0] + "." + sIdx[1]
					}
				}
			}
		})
		result = append(result, c)
	})
	return result, nil
}

// TODO: change order of return type
func parseIdx(chaptername string) (int, []string) {
	if strings.Contains(chaptername, "-") {
		re := regexp.MustCompile(`c(\d+)-(\d+)`)
		matches := re.FindStringSubmatch(chaptername)
		if len(matches) == 3 {
			return 0, matches[1:]
		}
		return -1, nil
	}
	chaptername = strings.ReplaceAll(chaptername, "prologue", "c0")
	re := regexp.MustCompile(`c(\d+)(?:\.(\d+)|)(?:\s+part(\d+)|)`)
	matches := re.FindStringSubmatch(chaptername)
	var rlmatch []string
	for i := 1; i < len(matches); i++ {
		if len(matches[i]) > 0 {
			rlmatch = append(rlmatch, matches[i])
		}
	}
	return 1, rlmatch
}

func getNextPage(doc *goquery.Document) (string, error) {
	np := doc.Find("a.next_page")
	if np.Length() > 0 {
		return np.First().AttrOr("href", ""), nil
	}
	return "", errors.New("next page not found")
}

// GetLatestSeries -
func GetLatestSeries() {
	// idx := 1
	// url := "https://www.novelupdates.com/latest-series/?st=1&pg=" + string(idx)
	// TODO: implement
}

func getNovelListing(doc *goquery.Document) map[string]string {
	novels := make(map[string]string)
	doc.Find("table#myTable > tbody > tr").Each(func(idx int, row *goquery.Selection) {
		col := row.Find("td:nth-child(3) > a")
		novels[col.Text()] = col.AttrOr("href", "")
	})
	return novels
}

// GetSeriesRanking - count
func GetSeriesRanking(count int) (map[string]string, error) {
	if count < 25 {
		return nil, errors.New("count has to be >= 1")
	}
	pages := (count + 25 - 1) / 25
	log.Printf("check %d/%d = %d\n", count, 25, pages)
	series := make(map[string]string)
	for i := 1; i <= pages; i++ {
		url := "https://www.novelupdates.com/series-ranking/?rank=popmonth&pg=" + string(i)
		doc, err := utils.GetSource(url)
		if err != nil {
			return series, err
		}
		novels := getNovelListing(doc)
		for k, v := range novels {
			series[k] = v
		}
	}
	return series, nil
}

// GetNovelsByPage - by idx
func GetNovelsByPage(idx int) (map[string]string, bool, error) {
	url := fmt.Sprintf("https://www.novelupdates.com/novelslisting/?st=%d", idx)
	doc, err := utils.GetSource(url)
	if err != nil {
		return nil, false, err
	}
	result := getNovelListing(doc)
	_, err = getNextPage(doc)
	return result, err != nil, nil
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
