package main

import "net/http"

func (a *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("GET /snippet/view/{id}", a.snippetView)
	mux.HandleFunc("GET /snippet/create", a.snippetCreate)
	mux.HandleFunc("POST /snippet/create", a.snippetCreatePost)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}
