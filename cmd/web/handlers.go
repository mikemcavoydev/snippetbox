package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	tmpls := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/navigation.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(tmpls...)
	if err != nil {
		a.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		a.serverError(w, r, err)
	}
}

func (a *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying specific snippet with ID: %d", id)
}

func (a *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func (a *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
