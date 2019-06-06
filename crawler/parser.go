package crawler

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
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
	config, err := p.GetConfig(url)
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

// GetConfig -- try to find config matching url
func (p *Parser) GetConfig(url string) (*lib.CrawlerConfig, error) {
	for _, config := range p.configs {
		if config.Match(url) {
			return &config, nil
		}
	}
	return nil, fmt.Errorf("no matching config found for %s", url)
}

// LoadCrawlers from plugins folder
func loadCrawlers() ([]lib.CrawlerConfig, error) {
	// create rel path
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("no caller information")
	}
	pluginsDir := path.Join(path.Dir(filename), "plugins")
	if !utils.Exists(pluginsDir) {
		return nil, errors.New("folder 'plugins' not found")
	}
	var configs []lib.CrawlerConfig
	err := filepath.Walk(pluginsDir, func(path string, info os.FileInfo, err error) error {
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

// GetConfigByID -- try to find config with id
func (p *Parser) GetConfigByID(id string) (*lib.CrawlerConfig, error) {
	for _, config := range p.configs {
		if config.GroupID == id {
			return &config, nil
		}
	}
	return nil, fmt.Errorf("couldn't find matching config for id: %s", id)
}
