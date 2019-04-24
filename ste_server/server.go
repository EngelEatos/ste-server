package ste_server

// wait for request from web & worker
// listen on ports: 1234, 1235(config?)
// handle requests on 

import (
	"net/rpc"
	"log"
	"net/http"
	"net"
)

type Queues struct {
	novelQueue []NovelQueueItem
	chapterQueue []ChapterQueueItem
}

type Args struct {
	Count int
}

var queues = new(Queues)

type ExQueues int

func (t *ExQueues) GetNovelQueue(args *Args, reply *[]NovelQueueItem) error {
	*reply = queues.novelQueue
	return nil
}

func (t *ExQueues) GetChapterQueue(args *Args, reply *[]ChapterQueueItem) error {
	*reply = queues.chapterQueue
	return nil
}

func (t *ExQueues) InsertNovel(novelitem *NovelQueueItem, reply *int) error {
	//insert
	queues.novelQueue = append(queues.novelQueue, *novelitem)
	*reply = 0
	return nil
} 

func (t *ExQueues) InsertChapter(chapteritem *ChapterQueueItem, reply *int) error {
	queues.chapterQueue = append(queues.chapterQueue, *chapteritem)
	*reply = 0
	return nil
}

func (t *ExQueues) PopNovel(args *Args, reply *NovelQueueItem) error {
	*reply, queues.novelQueue = queues.novelQueue[len(queues.novelQueue) - 1], queues.novelQueue[:len(queues.novelQueue) - 1]
	return nil
}

func (t *ExQueues) PopRangeNovel(count *int, reply *[]NovelQueueItem) error {
	*reply, queues.novelQueue = queues.novelQueue[len(queues.novelQueue) - *count:], queues.novelQueue[:len(queues.novelQueue) - *count]
	return nil
}

func (t *ExQueues) PopChapter(args *Args, reply *ChapterQueueItem) error {
	*reply, queues.chapterQueue = queues.chapterQueue[len(queues.chapterQueue) - 1], queues.chapterQueue[:len(queues.chapterQueue) - 1]
	return nil
}

func (t *ExQueues) PopRangeChapter(count *int, reply *[]ChapterQueueItem) error {
	*reply, queues.chapterQueue = queues.chapterQueue[len(queues.chapterQueue) - *count:], queues.chapterQueue[:len(queues.chapterQueue) - *count]
	return nil
}

func (t *ExQueues) NovelQueueLength(args *Args, reply *int) error {
	*reply = len(queues.novelQueue)
	return nil
}

func (t *ExQueues) ChapterQueueLength(args *Args, reply *int) error {
	*reply = len(queues.chapterQueue)
	return nil
}

func Run() {
	exQueues := new(ExQueues)
	rpc.Register(exQueues)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}