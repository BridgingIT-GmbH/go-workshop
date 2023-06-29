package repository

import (
	"fmt"
	"github.com/go-workshop/moviedb/model"
	"github.com/go-workshop/moviedb/service"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMovieDbRepository struct {
	client *mongo.Client
}

func NewMongoMovieDbRepository(client *mongo.Client) service.MovieDbRepository {
	return MongoMovieDbRepository{client}
}

func (repo MongoMovieDbRepository) FindAll() (*[]model.Movie, error) {
	return nil, fmt.Errorf("not implemented yet")
}
func (repo MongoMovieDbRepository) FindById(id uuid.UUID) (*model.Movie, error) {
	return nil, fmt.Errorf("not implemented yet")
}
func (repo MongoMovieDbRepository) CreateOrUpdate(movie model.Movie) (*model.Movie, error) {
	return nil, fmt.Errorf("not implemented yet")
}
func (repo MongoMovieDbRepository) Delete(id uuid.UUID) {

}
