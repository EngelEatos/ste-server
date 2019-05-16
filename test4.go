package main

import (
	"log"
	nuapi "ste/novelupdatesapi"
)

func main() {
	log.Println(nuapi.GetNovelsByPage(1))
}
