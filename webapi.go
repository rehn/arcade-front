package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func NewWebApi(c Configuration) *WebApi {
	var api WebApi
	api.config = c

	http.HandleFunc("/enable", api.enable)
	http.HandleFunc("/execute", api.exe)
	http.HandleFunc("/", api.HandleRequest)

	http.ListenAndServe(":"+strconv.Itoa(c.HttpPort), nil)
	return &api
}

type WebApi struct {
	config Configuration
}

func (api *WebApi) enable(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" || r.FormValue("value") == "" {
		http.Error(w, "Missing page or title data", http.StatusNotFound)
		return
	}
	value := (r.FormValue("value") == "true")
	db.UpdateEnabled(name, value)
}

func (api *WebApi) exe(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Missing name", http.StatusNotFound)
		return
	}
	game := db.GetGameData(name)
	execute <- game
}

func (api *WebApi) HandleRequest(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/" {
		p, err := PathToABS("assets/html/page.html")
		if api.NotFound(err, w, r) {
			return
		}
		t, err := template.New("page.html").ParseFiles(p)
		if api.NotFound(err, w, r) {
			return
		}

		data := api.buildData()

		err = t.Execute(w, data)
		api.NotFound(err, w, r)
	} else {
		p, err := PathToABS(strings.TrimLeft(r.RequestURI, "/"))
		if api.NotFound(err, w, r) {
			return
		}
		http.ServeFile(w, r, p)
	}
}

func (api *WebApi) buildData() ResponseData {
	data := ResponseData{}
	data.Config = api.config
	data.Games = db.GetGamesData("select * from items")
	return data
}

func (api *WebApi) NotFound(err error, w http.ResponseWriter, r *http.Request) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		LogWarning(err)
		return true
	}
	return false
}

type ResponseData struct {
	Config Configuration
	Games  []Game
}
