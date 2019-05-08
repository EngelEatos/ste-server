package updater

import "log"

const chapterlimit = 100
const novellimit = 10

type job struct {
}

type jobResult struct {
}

func workerNovel(job <-chan job, result <-chan jobResult) {

}

func workerChapter(job <-chan job, result <-chan jobResult) {
	// determine crawler
	// crawl chapter &Movim parse
	// save to Path
}

func getJobs(table string, count int) ([]job, error) {

	return []job{}, nil
}

func main() {
	// to worker
	cJobs, err := getJobs("chapter", chapterlimit)
	if err != nil {
		log.Fatal(err)
	}
	nJobs, err := getJobs("novel", novellimit)
	if err != nil {
		log.Fatal(err)
	}
	chapterJobs := make(chan job, len(cJobs))
	chapterResults := make(chan jobResult, len(nJobs))
	novelJobs := make(chan job, len(nJobs))
	novelResults := make(chan jobResult, len(cJobs))
	go workerNovel(novelJobs, novelResults)
	go workerChapter(chapterJobs, chapterResults)
	for _, e := range cJobs {
		chapterJobs <- e
	}
	close(chapterJobs)
	for _, e := range nJobs {
		novelJobs <- e
	}
	close(novelJobs)
	for i := 0; i < len(cJobs); i++ {
		<-chapterResults
	}
	close(chapterResults)
	for i := 0; i < len(nJobs); i++ {
		<-novelResults
	}
	close(novelResults)
}
