package main

import (
	"fmt"
	lib "ste/downloader/lib"
	"os"
	"path/filepath"
)

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

func main() {
	configs := LoadCrawlers()
	for _, crawler := range configs {
		crawler.Print()
	}
	fmt.Println("done")
}
