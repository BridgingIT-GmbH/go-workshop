package main

import (
	"fmt"
	"github.com/go-workshop/moviedb/rest"
	"log/slog"
	"net/http"
)

func main() {
	//inMemoryRepository := repository.NewInMemoryRepository()
	//TODO: add inMemoryRepository to controller constructor
	movieController := rest.NewMovieController()
	router := rest.NewRouter(movieController)

	port := 8080
	slog.Info(fmt.Sprintf("Starting server on port %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		slog.Error(err.Error())
	}
}
