package lib

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"regexp"
	"fmt"
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

func (c CrawlerConfig) Match(url string) bool {
	for _, regex := range c.UrlRegex {
		match, err := regexp.MatchString(regex, url)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			return true
		}
	}
	return false
}

type Config struct {
	Delay bool
	UserAgent string
}