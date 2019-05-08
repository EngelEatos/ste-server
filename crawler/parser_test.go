package crawler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	url := "https://www.wuxiaworld.com/novel/desolate-era"
	p := New()
	if len(p.configs) != 1 {
		t.Errorf("wrong length of configs. expected %d, go %d\n", 1, len(p.configs))
	}
	pr := p.Parse(url)

	if pr.NovelTitle != "Desolate Era" {
		t.Errorf("wrong novel tilte, expected Desolate Era, got %s", pr.NovelTitle)
	}
	expCoverURL := "https://cdn.wuxiaworld.com/images/covers/de.jpg?ver=6a8b21616e397749f1a334a49aafe82e09ef2db9"
	if pr.CoverURL != expCoverURL {
		t.Errorf("wrong cover url, expected %s, got %s", expCoverURL, pr.CoverURL)
	}

	if len(pr.ChapterMap) != 1451 {
		t.Errorf("wrong len of ChapterMap. expected %d, got %d", 1451, len(pr.ChapterMap))
	}
	b, err := json.MarshalIndent(pr.ChapterMap, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	t.Log(string(b))

}
