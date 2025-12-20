package routes

import (
	"Api-Aula1/controller"
	"Api-Aula1/handler"
	"Api-Aula1/middlewares"
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
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	return r
}

var rotasUsuarios = []Rota{
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controller.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controller.Login,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controller.GetAllUsers,
		RequerAutenticacao: true,
	},
}

var rotasLivros = []Rota{
	{
		URI:                "/books/search",
		Metodo:             http.MethodGet,
		Funcao:             handler.HandleSearch,
		RequerAutenticacao: true,
	},
	{
		URI:                "/books/{bookId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarLivro,
		RequerAutenticacao: true,
	},
	{
		URI:                "/books/{bookId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarLivro,
		RequerAutenticacao: true,
	},
}
