package main

import (
	"database/sql"
	"fmt"
	"log"
	na "ste/novelupdatesapi"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pandora"
	dbname   = "ste"
)

func main() {
	// novelupdatesapi.Start()
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	sql := "INSERT INTO %s (id, name) VALUES ($1, $2);"
	for k, v := range na.NovelTypes {
		csql := fmt.Sprintf(sql, "ste.novel_type")
		m, err := db.Exec(csql, v, k)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	}

	for k, v := range na.Languages {
		csql := fmt.Sprintf(sql, "ste.language")
		m, err := db.Exec(csql, v, k)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	}

	for k, v := range na.Genres {
		csql := fmt.Sprintf(sql, "ste.genre")
		m, err := db.Exec(csql, v, k)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	}

	for k, v := range na.Tags {
		csql := fmt.Sprintf(sql, "ste.tag")
		m, err := db.Exec(csql, v, k)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	}

	ccsql := "INSERT INTO ste.group (id, url, name) VALUES (1, 'https://www.wuxiaworld.com', 'wuxiaworld')"
	m, err := db.Exec(ccsql)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(m)

}
