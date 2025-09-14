package rest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type MovieController struct {
	repository MovieRepository
}

func NewMovieController(repository MovieRepository) *MovieController {
	return &MovieController{repository}
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
	var idAsString = request.PathValue("id")
	movie := controller.repository.FindById(idAsString)
	if movie == nil {
		http.Error(response, "movie not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(response).Encode(movie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
