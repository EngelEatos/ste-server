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

var connected = false
var conn *sql.DB

// Connect to db
func Connect() (*sql.DB, error) {
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
	connected = true
	return db, nil
}

// IsConnected return connected boolean
func IsConnected() bool {
	return connected
}

// TODO: insertNovel
// TODO: insertCover
