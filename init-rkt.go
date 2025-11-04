package main

import (
	"app/data"
	"app/handlers"
	"app/middleware"
	"log"
	"os"

	"github.com/m-goku/rkt"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init rkt
	rkt := &rkt.RKT{}
	err = rkt.New(path)
	if err != nil {
		log.Fatal(err)
	}

	rkt.AppName = "app"

	myMiddleware := &middleware.Middleware{
		App: rkt,
	}

	myHandlers := &handlers.Handlers{
		App: rkt,
	}

	app := &application{
		App:        rkt,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}
	app.App.Routes = app.routes()
	app.Models = data.NewModel(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
