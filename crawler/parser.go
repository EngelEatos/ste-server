package crawler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"ste-server/crawler/lib"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Parser - class parser
type Parser struct {
	configs []lib.CrawlerConfig
}

// ParseObj - class for Parser
type ParseObj struct {
	url       string
	selectors lib.Selector
	doc       *goquery.Document
}

// ParseResult - struct for parse result
type ParseResult struct {
	ChapterMap map[int]string
	NovelTitle string
	CoverURL   string
}

// New - Create new parser struct
func New() Parser {
	configs := LoadCrawlers()
	return Parser{configs}
}

// Parse - url
func (p Parser) Parse(url string) ParseResult {
	config, err := p.getConfig(url)
	if err != nil {
		log.Fatalf("no suitable conifg for url: %s", url)
	}
	doc := getSource(url)
	parseObj := ParseObj{url, config.Selectors, doc}
	chapterMap := parseObj.GetChapter()
	novelTitle := parseObj.GetTitle()
	cover := parseObj.GetCover()

	// var chapters []Chapter
	// for idx, chapterUrl := range(chapterMap) {
	// 	doc = getSource(chapterUrl)
	// 	parseObj = ParseObj(chapterUrl, config.Selectors, doc)

	// 	chapterTitle := parseObj.GetChapterTitle()
	// 	chapterBody := parseObj.GetChapterBody()
	// 	chapters = append(chappters, Chapter(chapterTitle, chapterBody, idx))
	// }

	return ParseResult{chapterMap, novelTitle, cover}
}

func (p Parser) getConfig(url string) (lib.CrawlerConfig, error) {
	for _, config := range p.configs {
		if config.Match(url) {
			return config, nil
		}
	}
	return lib.CrawlerConfig{}, fmt.Errorf("no matching config found for %s", url)
}

// NewParseObj - create new ParseObj
func NewParseObj(url string, selector lib.Selector, doc *goquery.Document) ParseObj {
	return ParseObj{url, selector, doc}
}

// GetChapter - getChapter
func (p ParseObj) GetChapter() map[int]string {
	current := 0
	result := make(map[int]string)
	p.doc.Find(p.selectors.Chapter).Each(func(index int, html *goquery.Selection) {
		text := html.Text()
		url, ok := html.Attr("href")
		if !ok {
			log.Printf("no url found for %d\n", current)
			url = ""
		}
		exp := regexp.MustCompile(`(\d+)`)
		match := exp.FindStringSubmatch(text)
		idx := -1
		if match != nil {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			idx = i
		} else {
			idx = current
			current++
		}
		result[idx] = url
	})
	return result
}

// GetCover - select cover url from doc
func (p ParseObj) GetCover() string {
	c := p.doc.Find(p.selectors.Cover).First()
	url, ok := c.Attr("href")
	if !ok {
		// err
		log.Fatal("no cover found")
	}
	return url
}

// GetTitle - select title from doc
func (p ParseObj) GetTitle() string {
	c := p.doc.Find(p.selectors.Title).First()
	return c.Text()
}

// GetChapterTitle - extract chapter title from chapter page
func (p ParseObj) GetChapterTitle() string {
	return p.doc.Find(p.selectors.ChapterTitle).First().Text()
}

// GetChapterBody - extract Chapter body from chapter page
func (p ParseObj) GetChapterBody() string {
	sl, err := p.doc.Find(p.selectors.ChapterBody).First().Html()
	if err != nil {
		log.Fatal(err)
	}
	return sl
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

// LoadCrawlers from plugins folder
func LoadCrawlers() []lib.CrawlerConfig {
	var configs []lib.CrawlerConfig
	err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		configs = append(configs, lib.LoadCrawlerConfig(path))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return configs
}
