package lib

import (
	"testing"
	"reflect"
)

func TestLoadCrawlerConfig(t *testing.T) {
	path := "../plugins/wuxiaworld.com.json"
	config := LoadCrawlerConfig(path)
	if reflect.DeepEqual(config, (CrawlerConfig{})) {
		t.Error("expected filled config, got empty")
	}
}

func TestMatch(t *testing.T) {
	path := "../plugins/wuxiaworld.com.json"
	config := LoadCrawlerConfig(path)
	url := `https://www.wuxiaworld.com/novel/test123$`
	if !config.Match(url) {
		t.Error("expected match, got false")
	}
}
