package ste_pool

const (
	chapterlimit 100
	novellimit 10
)

type job struct {

}

type jobResult struct {

}

func wokerNovel(job <-chan job, result <-chan jobResult) {

}

func workerChapter(job <-chan job, result <-chan jobResult) {
	// determine crawler
	// crawl chapter &Movim parse
	// save to Path
}

func getJobs(table, count) ([]job, error) {

	return new([]job), nil
}

func main() {
	// to worker
	cJobs := getJobs("chapter", chapterlimit)
	nJobs := getJobs("novel", novellimit)
	chapterJobs := make(chan job, len(cJobs))
	chapterResults := make(chan jobResult, len(nJobs))
	novelJobs := make(chan job, len(nJobs))
	novelResults := make(chan jobResult, len(cJobs))
	go workerNovel(novelJobs, novelResults)
	go workerChapter(chapterJobs, chapterResults)
	for _, e := range(cJobs) {
		chapterJobs <- e
	}
	close(chapterJobs)
	for _, e := range(nJobs) {
		novelJobs <- e
	}
	close(novelJobs)
	for i := 0 ; i < len(cJobs); i++ {
		<- chapterResults
	}
	close(chapterResults)
	for i := 0 ; i < len(nJobs); i++ {
		<- novelResults
	}
	close(novelResults)
}