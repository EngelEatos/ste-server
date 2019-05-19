package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// . "github.com/volatiletech/sqlboiler/queries/qm"
	"ste/models"
	nuapi "ste/novelupdatesapi"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pandora"
	dbname   = "ste"
)

// DBM Database Manager
type DBM struct {
	IsConnected bool
	DB          *sql.DB
}

// Connect to db
func (dbm *DBM) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Printf("connected to %s:%d\n", host, port)
	dbm.IsConnected = true
	dbm.DB = db
	return nil
}

// InsertNovel - ....
func (dbm *DBM) InsertNovel(ctx context.Context, novel *nuapi.Novel) (*models.Novel, error) {
	iCover, err := dbm.InsertCover(ctx, novel)
	inovel := &models.Novel{
		Title:               null.StringFrom(novel.Title),
		Chaptercount:        null.IntFrom(novel.Chaptercount),
		NovelIDSTR:          null.StringFrom(novel.NovelIDSTR),
		Type:                null.IntFrom(novel.Type),
		Description:         null.StringFrom(novel.Description),
		LanguageID:          null.IntFrom(novel.LanguageID),
		Year:                null.IntFrom(novel.Year),
		Status:              null.IntFrom(novel.Status),
		Licensed:            null.BoolFrom(novel.Licensed),
		CompletlyTranslated: null.BoolFrom(novel.CompletlyTranslated),
		CoverID:             null.IntFrom(iCover.ID),
	}
	err = inovel.Insert(ctx, dbm.DB, boil.Infer())
	if err != nil {
		log.Println(err)
		return inovel, err
	}
	return inovel, nil
}

// InsertCover insert cover into db, returns model.Cover and error if necessary
func (dbm *DBM) InsertCover(ctx context.Context, novel *nuapi.Novel) (*models.Cover, error) {
	iCover := &models.Cover{URL: novel.CoverURL, Downloaded: null.BoolFrom(false), Path: null.StringFrom("")}
	err := iCover.Insert(ctx, dbm.DB, boil.Infer())
	if err != nil {
		return nil, err
	}
	return iCover, nil
}
