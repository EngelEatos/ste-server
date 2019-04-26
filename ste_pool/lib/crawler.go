package lib

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Selector struct {
	Title string
	Cover string
	Chapter string
	ChapterTitle string
	ChapterBody string
}

type CrawlerConfig struct {
	PageUrl string
	Url string
	UrlRegex []string
	Names []string
	Selectors Selector
	KillTags []string
	RemoveTags []string
}

func readFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return dat
}

func LoadCrawlerConfig(path string) CrawlerConfig {
	data := readFile(path)
	var config CrawlerConfig
	if err := json.Unmarshal(data, &config); err != nil {
        panic(err)
    }
	return config
}

func (c CrawlerConfig) Print() {
	cjson, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cjson))
}

type Config struct {
	Delay bool
	UserAgent string
}

type Crawler struct {
	novelID string
	pageurl string
	selector map[string]string
}

func New(novelID string, pageurl string, selector map[string]string) Crawler {
	c := Crawler{novelID, pageurl, selector}
	return c
}

func (c Crawler) start() {
	c.update(map[string]string{})
}

func (c Crawler) update(current map[string]string) {

}

