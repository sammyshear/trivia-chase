package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/sammyshear/trivia-chase/views/pages"
)

func NewRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// pages and static assets
	indexPage := pages.Home()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/", templ.Handler(indexPage))

	// api routes
	mux.HandleFunc("POST /api/question", GetQuestion)
	mux.HandleFunc("GET /api/session", OpenSession)
	mux.HandleFunc("POST /api/answer", AnswerQuestion)

	return mux
}
