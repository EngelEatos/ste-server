package lib

type Chapter struct {
	Title string
	Body string
	Idx int
}


type Volume struct {
	Idx int
	Chapters []Chapter
	StartIdx int
	EndIdx int
}

type Novel struct {
	Title string
	Volumes []Volume
	CoverUrl string
}
