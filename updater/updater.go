package updater

import (
	"context"
	"database/sql"
	"log"
	"ste/models"
	nuapi "ste/novelupdatesapi"
	"ste/server"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

const chapterlimit = 100
const novellimit = 10

type job struct {
}

type jobResult struct {
}

func workerNovel(db *sql.DB, ctx context.Context, jobs <-chan *models.NovelQueue) {
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
		chapter, err := nuapi.GetAllChapter(n.NovelIDSTR.String)
		if err != nil {
			// insert into queue
			log.Println(err)
			continue
		}
	}

}

func workerChapter(job <-chan *models.ChapterQueue) {
	// determine crawler
	// crawl chapter &Movim parse
	// save to Path
	// trigger bookcreator
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
	go workerNovel(novelJobs)
	go workerChapter(chapterJobs)
	for _, e := range *cJobs {
		chapterJobs <- e
	}
	close(chapterJobs)
	for _, e := range *nJobs {
		novelJobs <- e
	}
	close(novelJobs)
}
