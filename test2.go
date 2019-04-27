package main

import (
	"fmt"
	"ste/novelupdatesAPI"
	"ste/server"
)

func main() {
	fmt.Println(novelupdatesAPI.NovelTypes)
	server.Webstart()
}
