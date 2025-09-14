package rest

import (
	"fmt"
	"log/slog"
	"net/http"
)

type MovieController struct {
	repository MovieRepository
}

func NewMovieController() *MovieController {
	return &MovieController{}
}

func (controller *MovieController) Hello(response http.ResponseWriter, request *http.Request) {
	slog.Info(fmt.Sprintf("%s %s", request.Method, request.URL))
	_, err := fmt.Fprint(response, "Hello!")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func (controller *MovieController) FindById(response http.ResponseWriter, request *http.Request) {
	slog.Info(fmt.Sprintf("%s %s", request.Method, request.URL))
	response.Header().Set("Content-Type", "application/json")

	//TODO
	//adjust routing (routes.go)
	//implement this method
	//fetch movie from InMemoryRepository by id

	//example for serializing a movie to json
	//json.NewEncoder(response).Encode(movie)
}
