package ste_server

import (
	// "time"
	"fmt"
)

type NovelQueueItem struct {
	NovelID int
	Title string
	QueuedAt string
}

func (i NovelQueueItem) String() string {
	return fmt.Sprintf("NovelID: %d, Title: %v, QueuedAt: %v\n",i.NovelID, i.Title, i.QueuedAt)
}

func (i NovelQueueItem) Print() {
	fmt.Printf(i.String())
}


type ChapterQueueItem struct {
	Title string
	Index int
	QueuedAt string
}

func (i ChapterQueueItem) String() string {
	return fmt.Sprintf("Title: %v, Index: %d, QueuedAt: %v\n", i.Title, i.Index, i.QueuedAt)
}

func (i ChapterQueueItem) Print() {
	fmt.Printf(i.String())
}

type Queues struct {
	novelQueue []NovelQueueItem
	chapterQueue []ChapterQueueItem
}

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