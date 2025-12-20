package controller

import (
	"Api-Aula1/autenticacao"
	"Api-Aula1/model"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"Api-Aula1/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := persistency.Connect()
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUsersRepo(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioSalvoNoBanco.Password, usuario.Password); erro != nil {
		responses.Err(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{Token: token})
}
