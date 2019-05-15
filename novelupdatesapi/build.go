package novelupdatesapi

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	// . "github.com/volatiletech/sqlboiler/queries/qm"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/queries"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "chaos"
	password = ""
	dbname   = "ste"
)

// ParseSQLFile - asdkljf
func ParseSQLFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	stringBuf := string(data)
	var result []string
	current := ""
	for _, line := range strings.Split(stringBuf, "\n") {
		line = strings.ReplaceAll(line, "\r", "")
		if !strings.HasPrefix(line, "-- ") && line != "" && line != "\n" {
			if strings.HasPrefix(line, "\t") {
				line = strings.Replace(line, "\t", "", 1)
			}
			current += " " + strings.Replace(line, "\t", "", 1)
			line = strings.TrimRight(line, "\r\n")
			if strings.HasSuffix(line, ";") {
				result = append(result, current)
				current = ""
			}
		}

	}
	return result
}

func createTables(db *sql.DB) {
	queriesQ := ParseSQLFile("ste_files/schema-ste.sql")
	for idx, query := range queriesQ {
		log.Printf("%d - %s", idx, query)
		err := queries.Raw(query, db, context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

}

// Start - ...
func Start() {
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", host, port, user)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	// create tables
 	createTables(db)
	// insert constants
	
	
}

