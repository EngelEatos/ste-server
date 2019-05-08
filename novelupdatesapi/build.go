package novelupdatesapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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
	fmt.Println(result)
	return result
}

func createTables() {
	queriesQ := ParseSQLFile("")
	// db, err := server.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	for idx, query := range queriesQ {
		log.Printf("%d - %s", idx, query)
		// queries.Raw(db, query)

	}

}

func start() {

}
