package controller

import (
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Err(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.User
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Prepare("cadastro"); erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepo(db)
	usuario.ID, erro = repositorio.Create(usuario)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepo(db)
	usuarios, erro := repositorio.FindAll()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Err(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.User
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Prepare("edicao"); erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepo(db)
	if erro = repositorio.Update(usuarioID, usuario); erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		responses.Err(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepo(db)
	if erro = repositorio.Delete(usuarioID); erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
