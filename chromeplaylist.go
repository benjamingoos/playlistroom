package main

import (
	"html/template"
	"net/http"
	"log"


	"test/chromeplaylist/youtube"

)

var tpl *template.Template



func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)
	http.HandleFunc("/videos", videos)
	http.HandleFunc("/player", player)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func search(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "search.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func videos(w http.ResponseWriter, req *http.Request) {
	searchField := req.FormValue("searchTerm")
	//log.Println(searchField)
	data := youtubeApi.YoutubeSearch(&searchField)
	err := tpl.ExecuteTemplate(w, "videos.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}

func player(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "player.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

