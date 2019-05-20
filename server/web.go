package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ste/models"
	nuapi "ste/novelupdatesapi"
	"strings"

	"github.com/oxtoacart/bpool"

	"github.com/gorilla/mux"
	"github.com/palantir/stacktrace"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type templateManager struct {
	templates *template.Template
	bufpool   *bpool.BufferPool
}

func newRouter(tpm *templateManager) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", tpm.searchHandler).Methods("GET")
	r.HandleFunc("/queues/", tpm.queuesHandler).Methods("GET")
	r.HandleFunc("/search/", tpm.searchHandler).Methods("GET")
	r.HandleFunc("/nu", nuHandler).Methods("POST")
	r.HandleFunc("/series/{serie}/", tpm.seriesHandler).Methods("GET")
	r.HandleFunc("/download/{serie}/", downloadHandler).Methods("GET")
	staticFileDirectory := http.Dir("./server/assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

// WebStart - start server
func WebStart() {
	// create templateManager
	tpm := new(templateManager)
	// load templates
	tpm.templates = template.Must(template.ParseGlob("server/templates/*"))
	// create buffer pool
	tpm.bufpool = bpool.NewBufferPool(64)
	// create router
	r := newRouter(tpm)
	// sart server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	novelid := vars["serie"]
	log.Printf("requested download for: %s\n", novelid)
}

func (tpm *templateManager) seriesHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println("novel doesnt exists in db")
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
		log.Printf("seriesHandler: %s\n", novel.Title)
		tpm.renderPage(w, r, "series", novel)
	} else {
		log.Printf("seriesHandler: %s\n", novel.Title)
		tpm.renderPage(w, r, "series", novel)
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

type terror struct {
	Title string
	Msg   string
}

func (tpm *templateManager) queuesHandler(w http.ResponseWriter, r *http.Request) {
	dbm := &DBM{}
	err := dbm.Connect()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	chapterQueue, err := models.ChapterQueues(qm.OrderBy("finished_at NULLS FIRST"), qm.Limit(100)).All(ctx, dbm.DB)
	if err != nil {
		log.Fatal(err)
	}
	novelQueue, err := models.NovelQueues(qm.OrderBy("finished_at NULLS FIRST"), qm.Limit(100)).All(ctx, dbm.DB)
	if err != nil {
		log.Fatal(err)
	}
	queues := queues{ChapterQueue: chapterQueue, NovelQueue: novelQueue}
	tpm.renderPage(w, r, "queues", queues)
}

func (tpm *templateManager) searchHandler(w http.ResponseWriter, r *http.Request) {
	tpm.renderPage(w, r, "home", nil)
}

func (tpm *templateManager) renderPage(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	buf := tpm.bufpool.Get()
	defer tpm.bufpool.Put(buf)
	err := tpm.templates.ExecuteTemplate(buf, name, data)
	if err != nil {
		log.Println(stacktrace.Propagate(err, "Can't load template"))
		w.WriteHeader(http.StatusInternalServerError)
		_ = tpm.templates.ExecuteTemplate(w, "error", terror{Title: "failed at " + r.URL.String(), Msg: "GTFO"})
		return
	}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	buf.WriteTo(w)
}
