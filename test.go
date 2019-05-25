package main

import (
	"log"
	nuapi "ste/novelupdatesapi"
)

func main() {

	chapters, err := nuapi.GetChapter("god-of-slaughter", 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("lenght: %d\n", len(chapters))
}
