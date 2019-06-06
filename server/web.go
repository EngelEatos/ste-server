package server

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ste/crawler"
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
	parser    *crawler.Parser
}

func newRouter(tpm *templateManager) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", tpm.searchHandler).Methods("GET")
	r.HandleFunc("/queues/", tpm.queuesHandler).Methods("GET")
	r.HandleFunc("/search/", tpm.searchHandler).Methods("GET")
	r.HandleFunc("/nu", nuHandler).Methods("POST")
	r.HandleFunc("/series/{serie}/", tpm.seriesHandler).Methods("GET")
	r.HandleFunc("/download/{serie}/", downloadHandler).Methods("GET")
	r.HandleFunc("/api", apiHandler).Methods("POST")
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
	t, err := template.ParseGlob("server/templates/*.html")
	if err != nil {
		log.Fatal("failed to load templates: ", err)
	}
	tpm.templates = template.Must(t, err)
	if err != nil {
		log.Fatal("failed to set templates: ", err)
	}
	// create buffer pool
	tpm.bufpool = bpool.NewBufferPool(64)
	// load CrawlerConfigs
	tpm.parser, err = crawler.New()
	if err != nil {
		log.Fatal("failed to create new Parser: ", err)
	}
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

type apiResponse struct {
	StatusCode int
	ErrorMsg   string
}

func (tpm *templateManager) apiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := apiResponse{StatusCode: 0, ErrorMsg: ""}
	// statusCode = 0 => OK ; statusCode != 0 => Errorg length
	// 				1 => wrong length
	//				2 => source not supported
	action := checkPostParam(w, 1, r.Form["action"])
	if len(action) == 0 {
		return
	}
	if action == "addsource" {
		url := checkPostParam(w, 1, r.Form["url"])
		if len(url) == 0 {
			return
		}
		log.Printf("[apiHandler]: {%s} - %s\n", action, url)
		// check if source is supported
		config, err := tpm.parser.GetConfig(url)
		if err != nil {
			// source not supported
			response.StatusCode = 2
			response.ErrorMsg = "source is not supported. Message admin or add source yourself!"
			respondWithJSON(w, response)
			return
		}
		// add supported source to novel

		respondWithJSON(w, response)
	}
}

func checkPostParam(w http.ResponseWriter, expectedLength int, param []string) string {
	response := apiResponse{StatusCode: 0, ErrorMsg: ""}
	if len(param) > expectedLength {
		errormsg := fmt.Sprintf("expected length of %d, got length of %d (%v)", expectedLength, len(param), param)
		log.Printf("%s\n", errormsg)
		response.StatusCode = 1
		response.ErrorMsg = errormsg
		respondWithJSON(w, response)
		return ""
	}
	return strings.Join(param[:expectedLength], "")
}

func respondWithJSON(w http.ResponseWriter, jsonData interface{}) {
	mjson, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(mjson)
}

type source struct {
	ChapterCount           int
	DownloadedChapterCount int
	Name                   string
	URL                    string
	Default                bool
}
type series struct {
	Novel   *models.Novel
	Sources []source
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
		novel, err := dbm.InsertNovel(pNovel)
		if err != nil {
			log.Fatal(err)
		}
		// pull chapter
		chapters, err := nuapi.GetAllChapter(pNovel.NovelIDSTR)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found %d chapters\n", len(chapters))
		for _, chapter := range chapters {
			// insert chapter into db
			c, err := dbm.InsertChapter(novel, &chapter)
			if err != nil {
				log.Fatal(err)
			}
			// insert chapter into queue
			_, err = dbm.InsertChapterQueue(novel, c)
			if err != nil {
				log.Fatal(err)
			}
		}
		// insert into novelqueue for next check
		_, err = dbm.InsertNovelQueue(novel)
		if err != nil {
			log.Fatal(err)
		}
		// get groups

		log.Printf("seriesHandler: %s\n", novel.Title)
		tpm.renderPage(w, r, "series.html", data{Active: "series", Data: series{Novel: novel, Sources: nil}})
	} else {
		log.Printf("seriesHandler: %s\n", novel.Title)
		tpm.renderPage(w, r, "series.html", data{Active: "series", Data: series{Novel: novel, Sources: nil}})
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
	tpm.renderPage(w, r, "queues.html", data{Data: queues, Active: "queues"})
}

type data struct {
	Data   interface{}
	Active string
}

func (tpm *templateManager) searchHandler(w http.ResponseWriter, r *http.Request) {
	tpm.renderPage(w, r, "home.html", data{Active: "home"})
}

func (tpm *templateManager) renderPage(w http.ResponseWriter, r *http.Request, name string, idata interface{}) {
	buf := tpm.bufpool.Get()
	defer tpm.bufpool.Put(buf)
	err := tpm.templates.ExecuteTemplate(buf, name, idata)
	if err != nil {
		log.Println(stacktrace.Propagate(err, "Can't load template"))
		w.WriteHeader(http.StatusInternalServerError)
		_ = tpm.templates.ExecuteTemplate(w, "error.html", data{Data: terror{Title: "failed at " + r.URL.String(), Msg: "GTFO"}, Active: ""})
		return
	}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	buf.WriteTo(w)
}
