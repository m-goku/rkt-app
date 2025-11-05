package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *Main) routes() *chi.Mux {
	//middleware must come before any routes
	//a.use(a.Middleware.CheckRemember)
	//add routes
	a.get("/", a.Handlers.Home)

	//static routes
	fileserver := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))
	return a.App.Routes
}
