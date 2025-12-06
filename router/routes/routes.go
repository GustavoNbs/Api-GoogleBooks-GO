package routes

import (
	"Api-Aula1/controller"
	"Api-Aula1/handler"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLivros...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

var rotasUsuarios = []Rota{
	{
		URI:    "/users",
		Metodo: http.MethodPost,
		Funcao: controller.CreateUser,
	},
	{
		URI:    "/users",
		Metodo: http.MethodGet,
		Funcao: controller.GetAllUsers,
	},
	{
		URI:    "/users/{id}",
		Metodo: http.MethodPut,
		Funcao: controller.UpdateUser,
	},
	{
		URI:    "/users/{id}",
		Metodo: http.MethodDelete,
		Funcao: controller.DeleteUser,
	},
}

var rotasLivros = []Rota{
	{
		URI:    "/books/search",
		Metodo: http.MethodGet,
		Funcao: handler.HandleSearch,
	},
}
