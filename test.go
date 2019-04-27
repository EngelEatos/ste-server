package main

import (
	"fmt"
	"log"
	"net/rpc"
	// "time"
	"ste/ste_server"
)

func main() {
	ste_server.Run()
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ni := &ste_server.NovelQueueItem{1, "test", "22.02.2000"}
	var reply int
	err = client.Call("ExQueues.InsertNovel", ni, &reply)
	if err != nil {
		log.Fatal("ExQueues error:", err)
	}
	ni.NovelID = 2
	ni.Title = "test2"
	err = client.Call("ExQueues.InsertNovel", ni, &reply)
	if err != nil {
		log.Fatal("ExQueues error:", err)
	}
	var reply2 []ste_server.NovelQueueItem
	args := &ste_server.Args{-1}
	err = client.Call("ExQueues.GetNovelQueue", args, &reply2)
	if err != nil {
		log.Fatal("ExQueues error:", err)
	}
	fmt.Printf("%#v\n", reply2)
	var reply3 []ste_server.NovelQueueItem
	var count int = 2
	err = client.Call("ExQueues.PopRangeNovel", &count, &reply3)
	fmt.Printf("%#v\n", reply3)
	if err != nil {
		log.Fatal("ExQueues error:", err)
	}
	err = client.Call("ExQueues.GetNovelQueue", args, &reply2)
	if err != nil {
		log.Fatal("ExQueues error:", err)
	}
	fmt.Printf("%#v\n", reply2)
}
