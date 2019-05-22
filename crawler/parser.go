package crawler

import (
	"fmt"
	"os"
	"path/filepath"
	"ste/crawler/lib"
	"ste/utils"
)

// Parser struct
type Parser struct {
	configs []lib.CrawlerConfig
}

// New create new parser
func New() (*Parser, error) {
	configs, err := loadCrawlers()
	if err != nil {
		return nil, err
	}
	return &Parser{configs}, nil
}

// Parse TODO
func (p *Parser) Parse(url string) (string, string, error) {
	config, err := p.getConfig(url)
	if err != nil {
		return "", "", err
	}
	doc, err := utils.GetSource(url)
	if err != nil {
		return "", "", err
	}

	title, err := doc.Find(config.Selectors.ChapterTitle).First().Html()
	if err != nil {
		return "", "", err
	}
	body, err := doc.Find(config.Selectors.ChapterBody).First().Html()
	if err != nil {
		return "", "", err
	}
	return title, body, err
}

func (p *Parser) getConfig(url string) (*lib.CrawlerConfig, error) {
	for _, config := range p.configs {
		if config.Match(url) {
			return &config, nil
		}
	}
	return nil, fmt.Errorf("no matching config found for %s", url)
}

// LoadCrawlers from plugins folder
func loadCrawlers() ([]lib.CrawlerConfig, error) {
	var configs []lib.CrawlerConfig
	err := filepath.Walk("./plugins", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		configs = append(configs, lib.LoadCrawlerConfig(path))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return configs, nil
}
