package service

import (
	"github.com/go-workshop/moviedb/model"
	"github.com/google/uuid"
)

type DomainMovieService struct {
	repo MovieRepository
}

func NewDomainMovieService(repo MovieRepository) *DomainMovieService {
	return &DomainMovieService{repo}
}

func (s *DomainMovieService) LoadAllMovies() (*[]model.Movie, error) {
	return s.repo.FindAll()
}

func (s *DomainMovieService) LoadMovieById(id string) (*model.Movie, error) {
	return s.repo.FindById(id)
}

func (s *DomainMovieService) CreateOrUpdateMovie(movie *model.Movie) (*model.Movie, error) {
	if &movie.Id == nil || movie.Id == "" {
		movie.Id = uuid.NewString()
	}
	return s.repo.CreateOrUpdate(movie)
}

func (s *DomainMovieService) DeleteMovie(id string) {
	s.repo.Delete(id)
}
