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
	movieRepo := repository.NewMongoRepository(mongoClient)

	movieService := service.NewDomainMovieService(movieRepo)
	movieController := rest.NewMovieController(movieService)

	log.Println(fmt.Sprintf("Server is ready at: '%v'", config.ServerHostPort))
	if err := http.ListenAndServe(config.ServerHostPort, rest.Routes(movieController)); err != nil {
		log.Fatal(err.Error())
	}
}
