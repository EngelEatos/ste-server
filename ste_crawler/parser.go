package ste_crawler

import (
	"log"
	"goquery"
	"os"
	"path/filepath"
	"fmt"
)

type Parser struct {
	configs []CrawlerConfig
}

type ParseObj struct {
	url string
	selectors Selector
	doc *goquery.Document
}

type ParseResult struct {
	ChapterMap map[int]string
	NovelTitle string
	CoverUrl string
}

func New() Parser {
	configs := LoadCrawlers()
	return Parser{configs}
}

func (p Parser) Parse(url string) ParseResult {
	config := getConfig(url)
	if !config {
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

func (p Parser) getConfig(url string) CrawlerConfig {
	for _,config := range p.configs {
		if config.Match(url) {
			return config
		}
	}
	return nil
}

func New(url string, config CrawlerConfig) {
	return ParseObj(url, config)
}

func parseInt(s string) int {
	
	return i
}
func (p ParseObj) GetChapter() map[int]string{
	current := 0
	result := make(map[int]string)
	p.doc.Find(p.selectors.Chapter).Each(func(index int, html *goquery.Selection) {
		text := html.Text()
		url, ok = html.Attr("href")
		if !ok {
			log.Printf("no url found for %d\n", current)
			url = ""
		}
		exp := regexp.MustCompile(`(\d+)`)
		match := exp.FindStringSubmatch(exp)
		idx := -1
		if match {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			idx = i
		} else {
			idx = current
			current += 1
		}
		result[idx] = url
	})
	return result
}

func (p ParseObject) GetCover() string {
	c := p.doc.Find(p.selectors.Cover).First()
	url, ok := c.Attr("href")
	if !ok {
		// err
		fmt.Fatalf("no cover found")
	}
	return url
}

func (p ParseObj) GetTitle() string {
	c := p.doc.Find(p.selectors.Title).First()
	return c.Text()
}

func (p ParseObj) GetChapterTitle() string {
	return p.doc.Find(p.selectors.ChapterTitle).First().Text()
}

func (p ParseObj) GetChapterBody() string {
	return p.doc.Find(p.selectors.ChapterBody).First()
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

func LoadCrawlers() []lib.CrawlerConfig {
	var configs []lib.CrawlerConfig
	err := filepath.Walk("./plugins", func (path string, info os.FileInfo, err error) error {
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