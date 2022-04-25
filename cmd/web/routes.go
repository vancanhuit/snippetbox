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

	r.With(app.sessionManager.LoadAndSave).Get("/", app.home)
	r.With(app.sessionManager.LoadAndSave).Get(
		"/snippet/view/{id}", app.snippetView)
	r.With(app.sessionManager.LoadAndSave).Get(
		"/snippet/create", app.snippetCreateView)
	r.With(app.sessionManager.LoadAndSave).Post(
		"/snippet/create", app.snippetCreatePost)

	return r
}
