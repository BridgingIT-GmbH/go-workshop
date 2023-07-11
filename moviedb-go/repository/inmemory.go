package repository

import (
	"github.com/go-workshop/moviedb/model"
	"golang.org/x/exp/maps"
)

type InMemoryRepository struct {
	data map[string]model.Movie
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[string]model.Movie)}
}

func (repo *InMemoryRepository) FindAll() (*[]model.Movie, error) {
	var values = maps.Values(repo.data)
	return &values, nil
}
func (repo *InMemoryRepository) FindById(id string) (*model.Movie, error) {
	movie := repo.data[id]
	return &movie, nil
}
func (repo *InMemoryRepository) CreateOrUpdate(movie *model.Movie) (*model.Movie, error) {
	repo.data[movie.Id] = *movie
	return movie, nil
}
func (repo *InMemoryRepository) Delete(id string) {
	delete(repo.data, id)
}
