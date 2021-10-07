package main

import (
	"net/http"
	"tpm/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func makeRoute() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.Home{}.Index)
	r.Get("/problems/new", controllers.Problems{}.New)
	r.Post("/problems/new", controllers.Problems{}.Create)
	r.Get("/problems/{problemId}", controllers.Problems{}.View)
	return r
}
