package controller

import (
	"Api-Aula1/autenticacao"
	"Api-Aula1/model"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SalvarLivro(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Err(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Err(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var book model.Book
	if erro = json.Unmarshal(corpoRequisicao, &book); erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	book.UserID = usuarioID
	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewBooksRepo(db)
	book.ID, erro = repositorio.Salvar(book)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, book)
}

func BuscarLivros(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Err(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewBooksRepo(db)
	livros, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, livros)
}

func AtualizarLivro(w http.ResponseWriter, r *http.Request) {
	usuarioID, _ := autenticacao.ExtrairUsuarioID(r)
	parametros := mux.Vars(r)
	bookID, _ := strconv.ParseUint(parametros["bookId"], 10, 64)

	corpoRequisicao, _ := io.ReadAll(r.Body)
	var book model.Book
	json.Unmarshal(corpoRequisicao, &book)

	db, _ := persistency.Connect()
	defer db.Close()

	repositorio := repository.NewBooksRepo(db)
	if erro := repositorio.Atualizar(bookID, usuarioID, book); erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarLivro(w http.ResponseWriter, r *http.Request) {
	usuarioID, _ := autenticacao.ExtrairUsuarioID(r)
	parametros := mux.Vars(r)
	bookID, _ := strconv.ParseUint(parametros["bookId"], 10, 64)

	db, _ := persistency.Connect()
	defer db.Close()

	repositorio := repository.NewBooksRepo(db)
	if erro := repositorio.Deletar(bookID, usuarioID); erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
