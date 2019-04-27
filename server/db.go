package server

import (
	"database/sql"
	"fmt"
	// . "github.com/volatiletech/sqlboiler/queries/qm"
	_ "github.com/lib/pq"
	"log"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "ste"
)

var connected bool = false
var conn *sql.DB

func Connect() (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
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
		return db,err
	}
	log.Printf("connected to %s:%d\n", host, port)
	connected = true
	return db, nil
}

// func Database() (*sql.DB, error) {
// 	if connectionInstance == nil {
// 		connectionInstance, err = connect()
// 		if err != nil {
// 			log.Println("connecting to db failed")
// 			return &sql.DB, err
// 		}
// 	}
// 	for c
// }
// IsConnected return connected boolean
func IsConnected() bool {
	return connected
}


