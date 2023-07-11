package repository

import (
	"context"
	"github.com/go-workshop/moviedb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	return &MongoRepository{client}
}

func (repo *MongoRepository) FindAll() (*[]model.Movie, error) {
	var movies []model.Movie
	cursor, err := withMovieCollection(repo.client).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}
	if movies == nil {
		return &[]model.Movie{}, nil
	} else {
		return &movies, nil
	}
}
func (repo *MongoRepository) FindById(id string) (*model.Movie, error) {
	filter := bson.D{{"id", id}}
	var movie model.Movie
	err := withMovieCollection(repo.client).FindOne(context.TODO(), filter).Decode(&movie)
	return &movie, err
}
func (repo *MongoRepository) CreateOrUpdate(movie *model.Movie) (*model.Movie, error) {
	filter := bson.D{{"id", movie.Id}}
	update := bson.M{"$set": movie}
	opts := options.Update().SetUpsert(true)
	_, err := withMovieCollection(repo.client).UpdateOne(context.TODO(), filter, update, opts)
	return movie, err
}
func (repo *MongoRepository) Delete(id string) {
	filter := bson.D{{"id", id}}
	_, err := withMovieCollection(repo.client).DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("coult not delete movie... %v", err.Error())
	}
}

func withMovieCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("moviedb").Collection("movie")
}
