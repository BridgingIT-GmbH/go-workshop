package rest

import (
	"github.com/go-chi/chi/v5"
)

func Routes(movieDbController *MovieController) *chi.Mux {

	r := chi.NewRouter()
	r.Get("/movies", movieDbController.FindAll)
	r.Get("/movies/{id}", movieDbController.FindById)
	r.Post("/movies", movieDbController.CreateOrUpdate)
	r.Delete("/movies/{id}", movieDbController.Delete)

	return r
}
