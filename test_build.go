package main

import (
	"ste/novelupdatesapi"
)

func main() {
	novelupdatesapi.ParseSQLFile("ste_files/schema-ste.sql")
}
