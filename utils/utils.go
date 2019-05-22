package utils

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// Padd string
func Padd(s string, wlen int, pchar string, left bool) string {
	if len(s) >= wlen {
		return s
	}
	c := wlen - len(s)
	padding := strings.Repeat(pchar, c)
	if left {
		return padding + s
	}
	return s + padding
}

// Zpadd string with 0
func Zpadd(s string, wlen int) string {
	return Padd(s, wlen, "0", true)
}

// ChunkString split string into chunks
func ChunkString(input string, chunksize int) []string {
	var result []string
	runes := bytes.Runes([]byte(input))
	for i := 0; i < len(runes); i += chunksize {
		eidx := i + chunksize
		if eidx > len(runes) {
			eidx = len(runes)
		}
		result = append(result, string(runes[i:eidx]))
	}
	return result
}

// GeneratePath from chapter id
func GeneratePath(outputPath string, id int) (string, error) {
	// 9223 3720 3685 4775 807
	// 0000 0000 0000 0000 0001/ 0000 0000 0000 0000 001.json
	padded := Zpadd(string(id), 20)
	chunked := ChunkString(padded, 4)
	p := strings.Join(chunked, "/")
	fullpath := path.Join(outputPath, p, fmt.Sprintf("%s.json", padded))
	err := os.MkdirAll(fullpath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return fullpath, nil
}

// FileExists checks if file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Strip \r\n from string
func Strip(s string) string {
	reg := regexp.MustCompile(`\r?\n`)
	s = reg.ReplaceAllString(s, "")
	return strings.Join(strings.Fields(s), "")
}

// ParseInt - parse int from string
func ParseInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return i, nil
}

// GetSource -- download source from url and put it into goquery
func GetSource(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// GetInt - return value from map[string]int
func GetInt(key string, dict map[string]int) (int, error) {
	if val, ok := dict[key]; ok {
		return val, nil
	}
	return -1, fmt.Errorf("%s not in %v", key, dict)
}
