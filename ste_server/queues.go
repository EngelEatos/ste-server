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