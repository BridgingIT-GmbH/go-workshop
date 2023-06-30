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

func (s MovieDbService) LoadMovieById(id string) (*model.Movie, error) {
	return s.repo.FindById(id)
}

func (s MovieDbService) CreateOrUpdateMovie(movie model.Movie) (*model.Movie, error) {
	if &movie.Id == nil || movie.Id == "" {
		movie.Id = uuid.NewString()
	}
	return s.repo.CreateOrUpdate(movie)
}

func (s MovieDbService) DeleteMovie(id string) {
	s.repo.Delete(id)
}
