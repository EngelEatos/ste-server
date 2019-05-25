package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// . "github.com/volatiletech/sqlboiler/queries/qm"
	"ste/models"
	nuapi "ste/novelupdatesapi"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pandora"
	dbname   = "ste"

	nextCheck = 24
)

// DBM Database Manager
type DBM struct {
	IsConnected bool
	DB          *sql.DB
	ctx         context.Context
}

// Connect to db
func (dbm *DBM) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	dbm.IsConnected = true
	dbm.DB = db
	dbm.ctx = context.Background()
	return nil
}

// InsertNovel - ....
func (dbm *DBM) InsertNovel(novel *nuapi.Novel) (*models.Novel, error) {
	iCover, err := dbm.InsertCover(novel)
	inovel := &models.Novel{
		Title:               novel.Title,
		Chaptercount:        null.IntFrom(novel.Chaptercount),
		NovelIDSTR:          novel.NovelIDSTR,
		NtypeID:             null.IntFrom(novel.Type),
		Description:         null.StringFrom(novel.Description),
		LanguageID:          null.IntFrom(novel.LanguageID),
		Year:                null.IntFrom(novel.Year),
		Status:              null.IntFrom(novel.Status),
		Licensed:            null.BoolFrom(novel.Licensed),
		CompletlyTranslated: null.BoolFrom(novel.CompletlyTranslated),
		CoverID:             null.IntFrom(iCover.ID),
		GroupID:             null.NewInt(-1, false),
		FetchedAt:           novel.FetchedAt,
	}
	err = inovel.Insert(dbm.ctx, dbm.DB, boil.Infer())
	if err != nil {
		log.Println(err)
		return inovel, err
	}
	return inovel, nil
}

// InsertCover insert cover into db, returns model.Cover and error if necessary
func (dbm *DBM) InsertCover(novel *nuapi.Novel) (*models.Cover, error) {
	iCover := &models.Cover{URL: novel.CoverURL, Downloaded: null.BoolFrom(false), Path: null.StringFrom("")}
	err := iCover.Insert(dbm.ctx, dbm.DB, boil.Infer())
	if err != nil {
		return nil, err
	}
	return iCover, nil
}

// InsertNovelQueue - insert novel into queue
func (dbm *DBM) InsertNovelQueue(novel *models.Novel) (*models.NovelQueue, error) {
	scheduledAt := time.Now().Add(time.Duration(nextCheck) * time.Hour)
	nq := &models.NovelQueue{
		NovelID:     novel.ID,
		QueuedAt:    time.Now(),
		Finished:    null.BoolFrom(false),
		ScheduledAt: scheduledAt,
	}
	err := nq.Insert(dbm.ctx, dbm.DB, boil.Infer())
	if err != nil {
		return nil, err
	}
	return nq, nil
}

// InsertChapter -- insert nuapi.Chapter into db
func (dbm *DBM) InsertChapter(novel *models.Novel, chapter *nuapi.Chapter) (*models.Chapter, error) {
	boil.DebugMode = true
	log.Printf("trying to insert %#v\n", chapter)
	ichapter := &models.Chapter{
		Title:      chapter.Title,
		URL:        chapter.URL,
		Idx:        chapter.Idx,
		Part:       null.IntFrom(chapter.Part),
		Downloaded: false,
	}
	// err := ichapter.Insert(dbm.ctx, dbm.DB, boil.Infer())
	// if err != nil {
	// 	return nil, err
	// }
	// add chapter to novel
	err := novel.AddChapters(dbm.ctx, dbm.DB, true, ichapter)
	if err != nil {
		ichapter.Delete(dbm.ctx, dbm.DB)
		return nil, err
	}
	return ichapter, nil
}

// InsertChapterQueue -- insert chapter of novel into queue
func (dbm *DBM) InsertChapterQueue(novel *models.Novel, chapter *models.Chapter) (*models.ChapterQueue, error) {
	boil.DebugMode = true
	iChapterQueue := &models.ChapterQueue{
		QueuedAt:  time.Now(),
		Finished:  null.BoolFrom(false),
		ChapterID: chapter.ID,
	}
	err := novel.AddChapterQueues(dbm.ctx, dbm.DB, true, iChapterQueue)
	// err := iChapterQueue.Insert()
	if err != nil {
		return nil, err
	}
	return iChapterQueue, nil
}
