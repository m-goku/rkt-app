package main

import (
	"app/data"
	"app/handlers"
	"app/middleware"

	"github.com/m-goku/rkt"
)

type application struct {
	App        *rkt.RKT
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	rocket := initApplication()
	rocket.App.ListenAndServe()
}
