package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		app.clientError(w, http.StatusMethodNotAllowed)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeaders)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Get("/", app.home)
	r.Get("/snippet/view/{id}", app.snippetView)
	r.Post("/snippet/create", app.snippetCreate)

	return r
}
