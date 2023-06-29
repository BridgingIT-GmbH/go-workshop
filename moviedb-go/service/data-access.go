package service

import (
	"github.com/go-workshop/moviedb/model"
	"github.com/google/uuid"
)

type MovieDbRepository interface {
	FindAll() (*[]model.Movie, error)
	FindById(id uuid.UUID) (*model.Movie, error)
	CreateOrUpdate(movie model.Movie) (*model.Movie, error)
	Delete(id uuid.UUID)
}
