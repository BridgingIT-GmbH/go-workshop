package rest

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-workshop/moviedb/model"
	"github.com/go-workshop/moviedb/service"
	"log"
	"net/http"
)

type MovieDbController struct {
	service service.MovieDbService
}

func NewMovieDbController(service service.MovieDbService) MovieDbController {
	return MovieDbController{service}
}
func (c MovieDbController) FindAll(w http.ResponseWriter, r *http.Request) {
	movies, err := c.service.LoadAllMovies()
	renderOrError(w, r, movies, err)
}
func (c MovieDbController) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movie, err := c.service.LoadMovieById(id)
	renderOrError(w, r, movie, err)
}
func (c MovieDbController) CreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie model.Movie
	err := decoder.Decode(&movie)
	if err == nil {
		movie, err2 := c.service.CreateOrUpdateMovie(movie)
		renderOrError(w, r, movie, err2)
	} else {
		renderOrError(w, r, movie, err)
	}
}
func (c MovieDbController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c.service.DeleteMovie(id)
}

func renderOrError(w http.ResponseWriter, r *http.Request, model interface{}, err error) {
	if err == nil {
		render.JSON(w, r, model)
	} else {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(nil)
	}
}
