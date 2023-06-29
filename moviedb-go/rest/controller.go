package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-workshop/moviedb/model"
	"github.com/go-workshop/moviedb/service"
	"net/http"
)

type MovieDbController struct {
	service service.MovieDbService
}

func NewMovieDbController(service service.MovieDbService) MovieDbController {
	return MovieDbController{service}
}
func (c MovieDbController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You called 'FindAll'"))
}
func (c MovieDbController) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("You called FindById with id: '%v'", id)))
}
func (c MovieDbController) CreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie model.Movie
	err := decoder.Decode(&movie)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("You called 'CreateOrUpdate' with wrong payload"))
		return
	}
	render.JSON(w, r, movie)
}
func (c MovieDbController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("You called Delete with id: '%v'", id)))
}
