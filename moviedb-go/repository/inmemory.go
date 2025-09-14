package repository

import (
	"github.com/go-workshop/moviedb/model"
	"maps"
	"slices"
)

type InMemoryRepository struct {
	data map[string]model.Movie
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: testData()}
}

func (repo *InMemoryRepository) FindById(id string) *model.Movie {
	movie, ok := repo.data[id]
	if ok {
		return &movie
	} else {
		return nil
	}
}

func (repo *InMemoryRepository) FindAll() []model.Movie {
	return slices.Collect(maps.Values(repo.data))
}

func (repo *InMemoryRepository) CreateOrUpdate(movie *model.Movie) *model.Movie {
	repo.data[movie.Id] = *movie
	return movie
}

func (repo *InMemoryRepository) Delete(id string) {
	delete(repo.data, id)
}
