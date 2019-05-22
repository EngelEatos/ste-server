package utils

import "testing"

func TestStrip(t *testing.T) {
	var cases = map[string]string{
		"hi       sup":     "hisup",
		"hi\nsup":          "hisup",
		"hi\r\nsup":        "hisup",
		"hi\tsup":          "hisup",
		"hi\n   sup\t\r\n": "hisup",
	}
	for o, result := range cases {
		if Strip(o) != result {
			t.Errorf("expected %s, but got %s", result, Strip(o))
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
		i, _ := ParseInt(o)
		if i != result {
			t.Errorf("expected %d, but got %d", result, i)
		}
	}
}
