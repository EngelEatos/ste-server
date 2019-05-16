package server

import (
	"database/sql"
	"fmt"
	"log"

	// . "github.com/volatiletech/sqlboiler/queries/qm"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
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
	conn        *sql.DB
}

// Connect to db
func (dbm *DBM) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(errors.Wrap(err, ""))
		return &sql.DB{}, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Println(errors.Wrap(err, ""))
		return db, err
	}
	log.Printf("connected to %s:%d\n", host, port)
	dbm.IsConnected = true
	dbm.conn = db
	return db, nil
}

// TODO: insertNovel
func (dbm *DBM) InsertNovel()

// TODO: insertCover
