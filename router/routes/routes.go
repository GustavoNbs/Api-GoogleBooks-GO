package routes

import (
	handler "Api-Aula1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
}
