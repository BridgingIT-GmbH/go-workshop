package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-workshop/moviedb/rest"
)

func Routes(movieDbController rest.MovieDbController) *chi.Mux {

	r := chi.NewRouter()
	r.Get("/movies", movieDbController.FindAll)
	r.Get("/movies/{id}", movieDbController.FindById)
	r.Post("/movies", movieDbController.CreateOrUpdate)
	r.Delete("/movies/{id}", movieDbController.Delete)

	return r
}
