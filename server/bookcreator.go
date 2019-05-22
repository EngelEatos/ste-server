package server

import (
	"context"
	"fmt"
	"github.com/bmaupin/go-epub"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"io/ioutil"
	"log"
	"ste/models"
	"ste/utils"
)

const (
	cssfile    = ""
	fontfile   = ""
	outputPath = "./raw_novel/"
)

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
		path, err := utils.GeneratePath(outputPath, chapter.ID)
		if err != nil {
			log.Fatal(err)
		}
		exists, err := utils.FileExists(path)
		if err != nil {
			log.Fatal(err)
		}
		if !exists {
			log.Printf("chapter file does not exist: %s\n", path)
			continue
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%04d.html", chapter.Idx)
		title := chapter.Title
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
