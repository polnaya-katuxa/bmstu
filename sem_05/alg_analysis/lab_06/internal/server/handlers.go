package server

import (
	"fmt"
	"html/template"
	"lab_06_01/internal/models"
	"lab_06_01/internal/query"
	"log"
	"net/http"
	"net/url"
)

type searchInfo struct {
	Query   string
	Results []models.Cat
	Error   string
	IsError bool
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func search(m *query.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		query := val.Get("text")

		cats, recErr := m.Recognize(query)
		if recErr != nil {
			log.Println(recErr)
		}

		ts, err := template.ParseFiles("./templates/search.tmpl")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, searchInfo{
			Query:   query,
			Results: cats,
			Error:   fmt.Sprint(recErr),
			IsError: recErr != nil,
		})
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}
}
