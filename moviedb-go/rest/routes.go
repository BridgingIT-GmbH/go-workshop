package rest

import (
	"net/http"
)

func NewRouter(controller *MovieController) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", controller.Hello)
	router.HandleFunc("GET /movies/{id}", controller.FindById)
	return router
}
