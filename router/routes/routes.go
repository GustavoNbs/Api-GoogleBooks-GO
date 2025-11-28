package routes

import (
	"Api-Aula1/controller"
	handler "Api-Aula1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
	r.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)
}
