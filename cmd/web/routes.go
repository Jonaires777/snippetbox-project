package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dinamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dinamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dinamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", dinamic.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", dinamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
