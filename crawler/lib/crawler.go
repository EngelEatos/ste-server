package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

// Selector -- struct to store css selectors
type Selector struct {
	Title        string
	Cover        string
	Chapter      string
	ChapterTitle string
	ChapterBody  string
}

// CrawlerConfig -- parsing Json-file
type CrawlerConfig struct {
	PageURL    string
	URL        string
	URLRegex   []string
	Names      []string
	Selectors  Selector
	KillTags   []string
	RemoveTags []string
	GroupID    string
}

func readFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return dat
}

// LoadCrawlerConfig -- load json-config file
func LoadCrawlerConfig(path string) CrawlerConfig {
	data := readFile(path)
	var config CrawlerConfig
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	config.GroupID = filename[0 : len(filename)-len(ext)]
	return config
}

// Print -- print config
func (c CrawlerConfig) Print() {
	cjson, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cjson))
}

// Match -- iterate and match configs with url
func (c CrawlerConfig) Match(url string) bool {
	for _, regex := range c.URLRegex {
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
	Delay     bool
	UserAgent string
}
