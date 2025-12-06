package router

import (
	"Api-Aula1/router/routes"

	"github.com/gorilla/mux"
)

// New retorna um novo router configurado
func New() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurar(r)
}
