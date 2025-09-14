package main

import (
	"fmt"
	"github.com/go-workshop/moviedb/repository"
	"github.com/go-workshop/moviedb/rest"
	"log/slog"
	"net/http"
)

func main() {
	movieRepository := repository.NewInMemoryRepository()
	movieController := rest.NewMovieController(movieRepository)
	router := rest.NewRouter(movieController)

	port := 8080
	slog.Info(fmt.Sprintf("Starting server on port %d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		slog.Error(err.Error())
	}
}
