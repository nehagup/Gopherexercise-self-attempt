package main

import (
	"encoding/json"
	"io/ioutil"
	"html/template"
	"net/http"
)

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text    string `json:"text"`
		Chapter string `json:"arc"`
	} `json:"options"`
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type Story map[string]Chapter

func main(){
	http.HandleFunc("/", foos)
	http.ListenAndServe(":8080", nil)
}

func foos(w http.ResponseWriter, r *http.Request)  {
	data := GetFile()
	if r.URL.Path == "/intro" || r.URL.Path == "/"{
		tpl.ExecuteTemplate(w, "index.html", data["intro"])
	} else {
		tpl.ExecuteTemplate(w, "index.html", data[r.URL.Path[1:]])
	}
}

func GetFile() map[string]Chapter{
	var data Story
	raw, _ := ioutil.ReadFile("gopher.json")
	json.Unmarshal(raw , &data)
	return data
}