package rest

import "github.com/go-workshop/moviedb/model"

type MovieService interface {
	LoadAllMovies() (*[]model.Movie, error)
	LoadMovieById(id string) (*model.Movie, error)
	CreateOrUpdateMovie(movie *model.Movie) (*model.Movie, error)
	DeleteMovie(id string)
}
