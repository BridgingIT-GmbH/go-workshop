package main

import (
	"fmt"
	"github.com/go-workshop/moviedb/config"
	"github.com/go-workshop/moviedb/repository"
	"github.com/go-workshop/moviedb/rest"
	"github.com/go-workshop/moviedb/service"
	"log"
	"net/http"
)

func main() {
	mongoClient := config.CreateClient(config.MongoUri)
	movieDbRepo := repository.NewMongoRepository(mongoClient)

	movieDbService := service.NewDomainMovieService(movieDbRepo)
	movieDbController := rest.NewMovieController(movieDbService)

	log.Println(fmt.Sprintf("Server is ready at: '%v'", config.ServerHostPort))
	if err := http.ListenAndServe(config.ServerHostPort, rest.Routes(movieDbController)); err != nil {
		log.Fatal(err.Error())
	}
}
