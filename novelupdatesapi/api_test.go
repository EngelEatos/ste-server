package novelupdatesapi

import (
	"reflect"
	"testing"
)

func TestGetBool(t *testing.T) {
	var cases = map[string]bool{
		"No":  false,
		"Yes": true,
		"no":  false,
		"yes": true,
		"nO":  false,
		"yEs": true,
	}
	for o, result := range cases {
		if getBool(o) != result {
			t.Errorf("expected %t, but got %t", result, getBool(o))
		}
	}
}

func TestStrip(t *testing.T) {
	var cases = map[string]string{
		"hi       sup":     "hisup",
		"hi\nsup":          "hisup",
		"hi\r\nsup":        "hisup",
		"hi\tsup":          "hisup",
		"hi\n   sup\t\r\n": "hisup",
	}
	for o, result := range cases {
		if strip(o) != result {
			t.Errorf("expected %s, but got %s", result, strip(o))
		}
	}
}

func TestParseInt(t *testing.T) {
	var cases = map[string]int{
		"1":    1,
		"10":   10,
		"2222": 2222,
	}
	for o, result := range cases {
		if parseInt(o) != result {
			t.Errorf("expected %d, but got %d", result, parseInt(o))
		}
	}
}

func TestParseNovel(t *testing.T) {
	title := "The Human Emperor"
	chaptercount := 1782
	NovelIDSTR := "the-human-emperor"
	Type := 2444
	Genres := []int{8, 9, 5, 14, 3954}
	Tags := []int{7909, 306, 9522, 111, 9907, 782, 5708, 5944, 14593, 21268, 15026, 93, 192, 14431, 10134, 12035, 6526, 3257, 589, 2666, 265, 3019, 298, 8801, 1124, 3314, 334, 24959, 2571, 76, 4777, 7357, 148, 12907, 101, 71}
	LanguageID := 495
	Year := 2016
	Status := 0
	Licensed := true
	CompletlyTranslated := false
	Recommendations := map[string]string{"A Cheeky Kendo God": "https://www.novelupdates.com/series/a-cheeky-kendo-god/", "I am the Monarch": "https://www.novelupdates.com/series/i-am-the-monarch/", "Reverend Insanity": "https://www.novelupdates.com/series/reverend-insanity/", "The Second Coming of Avarice": "https://www.novelupdates.com/series/the-second-coming-of-avarice/", "Transcending the Nine Heavens": "https://www.novelupdates.com/series/transcending-the-nine-heavens/", "Zhanxian": "https://www.novelupdates.com/series/zhanxian/"}
	Authors := []string{"Huangfu Qi", "皇甫奇"}
	novel := ParseNovel("the-human-emperor")
	if title != novel.Title {
		t.Errorf("expected %s, got %s", title, novel.Title)
	}
	if chaptercount != novel.Chaptercount {
		t.Errorf("expected %d, got %d", chaptercount, novel.Chaptercount)
	}

	if NovelIDSTR != novel.NovelIDSTR {
		t.Errorf("expected %s, got %s", NovelIDSTR, novel.NovelIDSTR)
	}

	if Type != novel.Type {
		t.Errorf("expected %d, got %d", Type, novel.Type)
	}

	if !reflect.DeepEqual(Genres, novel.Genres) {
		t.Errorf("expected %v, got %v", Genres, novel.Genres)
	}

	if !reflect.DeepEqual(Tags, novel.Tags) {
		t.Errorf("expected %v, got %v", Tags, novel.Tags)
	}

	if LanguageID != novel.LanguageID {
		t.Errorf("expected %d, got %d", LanguageID, novel.LanguageID)
	}

	if Year != novel.Year {
		t.Errorf("expected %d, got %d", Year, novel.Year)
	}

	if Status != novel.Status {
		t.Errorf("expected %d, got %d", Status, novel.Status)
	}

	if Licensed != novel.Licensed {
		t.Errorf("expected %t, got %t", Licensed, novel.Licensed)
	}

	if CompletlyTranslated != novel.CompletlyTranslated {
		t.Errorf("expected %t, got %t", CompletlyTranslated, novel.CompletlyTranslated)
	}

	if !reflect.DeepEqual(Recommendations, novel.Recommendations) {
		t.Errorf("expected %v, got %v", Recommendations, novel.Recommendations)
	}

	if !reflect.DeepEqual(Authors, novel.Authors) {
		t.Errorf("expected %v, got %v", Authors, novel.Authors)
	}
}
