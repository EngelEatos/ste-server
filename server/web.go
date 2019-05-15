package server

import (
	"context"
	// "encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ste/models"
	"strconv"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/queries/qm"
	nuapi "ste/novelupdatesapi"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/", func(w *http.R))
	r.HandleFunc("/request", requestHandler).Methods("GET")
	r.HandleFunc("/submit", submitHandler).Methods("POST")
	r.HandleFunc("/queues", queuesHandler).Methods("GET")
	r.HandleFunc("/search/", searchHandler).Methods("GET")
	r.HandleFunc("/api/", apiHandler).Methods("POST")
	r.HandleFunc("/nu", nuHandler).Methods("POST")
	// live search
	// result
	staticFileDirectory := http.Dir("./server/assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

// WebStart - start server
func WebStart() {
	r := newRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

}

func nuHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("searchTerm")
	result := nuapi.LiveSearch(searchTerm)
	fmt.Fprintf(w, result[:len(result)-1])
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		log.Fatal("couldnt connect to db")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	sources, err := models.Sources().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("server/templates/request.gohtml")
	t.Execute(w, sources)
}

type queues struct {
	ChapterQueue []*models.ChapterQueue
	NovelQueue   []*models.NovelQueue
}

func queuesHandler(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		log.Fatal("couldnt connect to db")
	}
	chapterQueue, err := models.ChapterQueues(qm.OrderBy("finishedAt NULLS FIRST"), qm.Limit(100)).All(context.Background(), db)
	novelQueue, err := models.NovelQueues(qm.OrderBy("finishedAt NULLS FIRST"), qm.Limit(100)).All(context.Background(), db)
	queues := queues{ChapterQueue: chapterQueue, NovelQueue: novelQueue}
	t, _ := template.ParseFiles("server/templates/queues.gohtml")
	t.Execute(w, queues)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// db, err := Connect()
	// if err != nil {
	// 	log.Fatal("couldnt connect to db")
	// }
	// searchterm := `%` + r.FormValue("term") + `%`
	// if searchterm == `%%` {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(request{})
	// } else {
	// 	novels, err := models.Novels(qm.Where("title like ?", searchterm), qm.Or("description like ?", searchterm), qm.Or("novel_id_str like ?", searchterm), qm.Limit(5)).All(context.Background(), db)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(novels)
	// }
	// t, _ := template.ParseFiles("server/templates/search.html")
	dat, err := ioutil.ReadFile("server/templates/search.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(dat))
}

type request struct {
	URL   string
	Title string
	Src   int
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	title := r.FormValue("title")
	src := r.FormValue("src")
	fmt.Println(url, title, src)
	t, _ := template.ParseFiles("server/templates/submit.gohtml")
	srcInt, err := strconv.ParseInt(src, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	req := request{URL: url, Title: title, Src: int(srcInt)}
	t.Execute(w, &req)
	// check if valid else contact admin
	// insert new novel in db
	// insert novel in queue
}
