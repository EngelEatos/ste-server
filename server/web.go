package server

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"ste/models"
	nuapi "ste/novelupdatesapi"
	"strings"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/queues", queuesHandler).Methods("GET")
	r.HandleFunc("/search/", searchHandler).Methods("GET")
	r.HandleFunc("/nu", nuHandler).Methods("POST")
	r.HandleFunc("/series/{serie}/", seriesHandler).Methods("GET")
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

func seriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	novelid := vars["serie"]
	// check if already in db
	dbm := &DBM{}
	err := dbm.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	novel, err := models.Novels(qm.Where("novel_id_str=?", novelid)).One(ctx, dbm.DB)
	if err != nil {
		fmt.Println("novel doesnt exists in db")
		// parse now
		pNovel, err := nuapi.ParseNovel(novelid)
		if err != nil {
			log.Fatal(err)
		}
		// insert novel and cover
		novel, err := dbm.InsertNovel(ctx, pNovel)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(novel.Title.String)
		t, _ := template.ParseFiles("server/templates/series.html")
		t.Execute(w, novel)
	} else {
		fmt.Println(novel.Title)
		t, _ := template.ParseFiles("server/templates/series.html")
		t.Execute(w, novel)
	}
}

func nuHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("searchTerm")
	result, err := nuapi.LiveSearch(searchTerm)
	if err != nil {
		log.Fatal(err)
	}
	result = strings.ReplaceAll(result, "https://www.novelupdates.com/", "../")
	fmt.Fprintf(w, result[:len(result)-1])
}

type queues struct {
	ChapterQueue []*models.ChapterQueue
	NovelQueue   []*models.NovelQueue
}

func queuesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: implement

	// db, err := Connect()
	// if err != nil {
	// 	log.Fatal("couldnt connect to db")
	// }
	// chapterQueue, err := models.ChapterQueues(qm.OrderBy("finishedAt NULLS FIRST"), qm.Limit(100)).All(context.Background(), db)
	// novelQueue, err := models.NovelQueues(qm.OrderBy("finishedAt NULLS FIRST"), qm.Limit(100)).All(context.Background(), db)
	// queues := queues{ChapterQueue: chapterQueue, NovelQueue: novelQueue}
	// t, _ := template.ParseFiles("server/templates/queues.gohtml")
	// t.Execute(w, queues)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("server/templates/search.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(dat))
}
