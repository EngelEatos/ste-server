package updater

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"ste/crawler"
	"ste/models"
	nuapi "ste/novelupdatesapi"
	"ste/server"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	chapterlimit = 100
	novellimit   = 10
	outputPath   = "./raw_novel/"
)

// TODO: implement updateNovel
// func updateNovel(ctx context.Context, db *sql.DB, n *models.Novel, novel *nuapi.Novel) *models.Novel {
// 	n.Title = novel.Title
// 	n.Chaptercount = novel.Chaptercount
// 	n.NovelIDSTR = novel.NovelIDSTR
// 	ntype, err := models.FindNovelType(ctx, db, novel.Type)
// 	if err != nil {
// 		return n, fmt.Errorf("couldnt find noveltype: %d\n", novel.Type)
// 	}
// 	n.SetNtype(ctx, db, true, ntype)
// 	// n.setG
// }

// Difference TODO:
func Difference(a models.ChapterSlice, b []nuapi.Chapter) (diff []nuapi.Chapter) {
	m := make(map[int]bool)
	for _, item := range a {
		m[item.Idx.Int] = true
	}

	for _, item := range b {
		if _, ok := m[item.Idx]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func workerNovel(ctx context.Context, db *sql.DB, jobs <-chan *models.NovelQueue) {
	// parse novel page and get chapters
	// insert chapters into chapterqueue
	for j := range jobs {
		n, err := j.Novel().One(ctx, db)
		if err != nil {
			log.Println(err)
			continue
		}
		novel, err := nuapi.ParseNovel(n.NovelIDSTR.String)
		if err != nil {
			log.Println(err)
			continue
		}
		// upadate novel
		dbNovel, err := models.Novels(qm.Where("novel_id_str=?", novel.NovelIDSTR)).One(ctx, db)
		if err != nil {
			log.Println(err)
			continue
		}
		dbNovel.Update(ctx, db, boil.Infer())
		chapter, err := nuapi.GetAllChapter(n.NovelIDSTR.String)
		if err != nil {
			log.Println(err)
			continue
		}
		dbChapter, err := dbNovel.Chapters(qm.OrderBy("idx ASC")).All(ctx, db)
		if err != nil {

		}
		// TODO: compare dbChapter and chapter
		for a, k := range dbChapter {

		}
		// TODO: difference between chapter and chapter in db

		// TODO: insert into queue with next scheduledTime

	}

}

// Chapter ..
type Chapter struct {
	Title string
	Body  string
}

func workerChapter(ctx context.Context, db *sql.DB, jobs <-chan *models.ChapterQueue) {
	for j := range jobs {
		// url
		chapter, err := j.Chapter().One(ctx, db)
		if err != nil {
			log.Printf("failed to get Chapter for ID %d: %s\n", chapter.ID, err)
			continue
		}
		url := chapter.URL.String
		// determine crawler
		parser, err := crawler.New()
		if err != nil {
			log.Printf("failed to create new Crawler: %s\n", err)
			continue
		}
		// crawl chapter
		title, body, err := parser.Parse(url)
		if err != nil {
			log.Printf("failed to parse url %s: %s\n", url, err)
			continue
		}
		// generate Path
		oPath, err := utils.generatePath(outputPath, chapter.ID)
		if err != nil {
			log.Printf("failed to generate path for %d: %s\n", chapter.ID, err)
			continue
		}
		data, err := json.MarshalIndent(Chapter{title, body}, "", " ")
		if err != nil {
			log.Printf("failed to create json: %s\n", err)
			continue
		}
		err = ioutil.WriteFile(oPath, data, 0644)
		if err != nil {
			log.Printf("failed to write json-data to file: %s\n", err)
			continue
		}
		chapter.Downloaded = null.BoolFrom(true)
		// TODO: trigger bookcreator??
		// update chapter in queue
		j.Finished = null.BoolFrom(true)
		j.FinishedAt = null.TimeFrom(time.Now())
		_, err = j.Update(ctx, db, boil.Infer())
		if err != nil {
			log.Printf("failed to set job %d to finish: %s\n", chapter.ID, err)
			continue
		}
	}
}

func getChapterJobs(count int) (*models.ChapterQueueSlice, error) {
	dbm := &server.DBM{}
	dbm.Connect()
	ctx := context.Background()
	chapterJobs, err := models.ChapterQueues(qm.Limit(count)).All(ctx, dbm.DB)
	if err != nil {
		return nil, err
	}
	return &chapterJobs, nil
}

func getNovelJobs(count int) (*models.NovelQueueSlice, error) {
	dbm := &server.DBM{}
	dbm.Connect()
	ctx := context.Background()
	novelJobs, err := models.NovelQueues(qm.Limit(count)).All(ctx, dbm.DB)
	if err != nil {
		return nil, err
	}
	return &novelJobs, nil
}

func main() {
	// to worker
	dbm := &server.DBM{}
	dbm.Connect()
	ctx := context.Background()
	cJobs, err := getChapterJobs(chapterlimit)
	if err != nil {
		log.Fatal(err)
	}
	nJobs, err := getNovelJobs(novellimit)
	if err != nil {
		log.Fatal(err)
	}
	chapterJobs := make(chan *models.ChapterQueue, chapterlimit)
	novelJobs := make(chan *models.NovelQueue, novellimit)
	go workerNovel(ctx, dbm.DB, novelJobs)
	go workerChapter(ctx, dbm.DB, chapterJobs)
	for _, e := range *cJobs {
		chapterJobs <- e
	}
	close(chapterJobs)
	for _, e := range *nJobs {
		novelJobs <- e
	}
	close(novelJobs)
}
