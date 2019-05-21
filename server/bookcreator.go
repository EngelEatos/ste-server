package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"ste/models"

	"github.com/bmaupin/go-epub"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	cssfile  = ""
	fontfile = ""
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// func rpadd(s string, wlen int, pchar string) string {
// 	return padd(s, wlen, pchar, false)
// }

// func lpadd(s string, wlen int, pchar string) string {
// 	return padd(s, wlen, pchar, true)
// }

// func zpadd(s string, wlen int) string {
// 	return padd(s, wlen, "0", true)
// }
// i=0; 0 - 99
// i=1; 100 - 199
// i=2; 200 - 299
// x;   x*100 - (x+1)*100 - 1
func sliceNdice(chapters []models.Chapter) [][]models.Chapter {
	vCount := (len(chapters) + 100 - 1) / 100
	result := make([][]models.Chapter, vCount)
	for i := 0; i < vCount; i++ {
		eIdx := (i + 1) * 100
		if eIdx > len(chapters) {
			eIdx = len(chapters)
		}
		result[i] = chapters[i*100 : eIdx]
	}
	return result
}

// Build as parameter []Chapter
func Build(novel *models.Novel, oPath string) error {
	dbm := &DBM{}
	dbm.Connect()
	ctx := context.Background()
	chapters, err := novel.Chapters(qm.OrderBy("idx ASC")).All(ctx, dbm.DB)
	if err != nil {
		return err
	}

	e := epub.NewEpub(novel.Title)
	cover, err := novel.Cover().One(ctx, dbm.DB)
	if err != nil {
		log.Fatal(err)
	}
	e.SetCover(cover.Path.String, "")

	csspath, err := e.AddCSS(cssfile, "")
	if err != nil {
		log.Fatal(err)
	}

	_, err = e.AddFont(fontfile, "")
	if err != nil {
		log.Fatal(err)
	}

	for _, chapter := range chapters {
		if !chapter.Path.Valid || !fileExists(chapter.Path.String) {
			log.Printf("chapter file does not exist: %s\n", chapter.Path.String)
			continue
		}
		data, err := ioutil.ReadFile(chapter.Path.String)
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%04d.html", chapter.Idx.Int)
		title := chapter.Title.String
		if !chapter.Title.Valid {
			title = fmt.Sprintf("Chapter %d", chapter.Idx.Int)
		}
		_, err = e.AddSection(string(data), title, filename, csspath)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = e.Write(oPath)
	if err != nil {
		return err
	}
	return nil
}
