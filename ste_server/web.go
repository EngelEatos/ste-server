package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"log"
)

type Page struct {
    Title string
    Body  []byte
}

func main() {
    http.HandleFunc("/queues/", queuesHandler)
    http.HandleFunc("/request/", requestHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "request", nil)
}

func queuesHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "queues", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    title := r.FormValue("title")
    
}
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles("templates/" + tmpl + ".html")
    t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}