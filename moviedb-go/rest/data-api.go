package rest

import (
	"github.com/go-workshop/moviedb/model"
)

type MovieRepository interface {
	FindById(id string) *model.Movie
}
