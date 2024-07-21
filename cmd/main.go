package cmd

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/sammyshear/trivia-chase/views/pages"
)

func App() {
	mux := http.NewServeMux()
	indexPage := pages.Home()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/", templ.Handler(indexPage))

	http.ListenAndServe(":3000", mux)
}
