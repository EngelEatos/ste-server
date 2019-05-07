package novelupdatesapi

import (
	"io/ioutil"
	"log"
	"ste/server"
	"strings"

	"github.com/volatiletech/sqlboiler/queries"
)

// ParseSQLFile - asdkljf
func ParseSQLFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	stringBuf := string(data)
	var queries []string
	current := ""
	for _, line := range strings.Split(stringBuf, "\n") {
		if !strings.HasPrefix(line, "-- ") && line != "" && line != "\n" {
			current += " " + strings.Replace(line, "\t", "", 1)
			if strings.HasSuffix(line, ";") {
				queries = append(queries, current)
				current = ""
			}
		}
	}
	return queries
}

func createTables() {
	queriesQ := ParseSQLFile("")
	db, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}
	for idx, query := range queriesQ {
		log.Printf("%d - %s", idx, query)
		queries.Raw(db, query)

	}

}

func start() {

}
