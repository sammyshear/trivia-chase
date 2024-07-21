package cmd

import (
	"net/http"

	"github.com/sammyshear/trivia-chase/api"
)

func App() {
	mux := api.NewRoutes()
	http.ListenAndServe(":3000", mux)
}
