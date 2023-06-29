package service

import (
	"github.com/go-workshop/moviedb/model"
	"github.com/google/uuid"
)

type MovieDbService struct {
	repo MovieDbRepository
}

func NewMovieDbService(repo MovieDbRepository) MovieDbService {
	return MovieDbService{repo}
}

func (s MovieDbService) LoadAllMovies() (*[]model.Movie, error) {
	return s.repo.FindAll()
}

func (s MovieDbService) LoadMovieById(id uuid.UUID) (*model.Movie, error) {
	return s.repo.FindById(id)
}

func (s MovieDbService) CreateOrUpdateMovie(movie model.Movie) (*model.Movie, error) {
	return s.repo.CreateOrUpdate(movie)
}

func (s MovieDbService) DeleteMovie(id uuid.UUID) {
	s.repo.Delete(id)
}
