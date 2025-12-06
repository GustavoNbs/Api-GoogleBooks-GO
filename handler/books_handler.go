package handler

import (
	"Api-Aula1/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RespostaDoGoogle struct {
	Items []struct {
		VolumeInfo DadosDoLivro `json:"volumeInfo"`
	} `json:"items"`
}
type DadosDoLivro struct {
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
	PageCount   int      `json:"pageCount"`
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebi uma requisição de busca de livros!")

	nomePraPesquisar := r.URL.Query().Get("nome")
	if nomePraPesquisar == "" {
		responses.Err(w, http.StatusBadRequest, fmt.Errorf("parametro 'nome' é obrigatório"))
		return
	}

	livros, erro := buscarNoGoogle(nomePraPesquisar)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, livros)
}

// buscarNoGoogle faz a requisição externa à API e trata os dados
func buscarNoGoogle(nome string) ([]interface{}, error) {
	nomeFormatado := url.QueryEscape(nome)
	urlGoogle := "https://www.googleapis.com/books/v1/volumes?q=" + nomeFormatado

	resp, erro := http.Get(urlGoogle)
	if erro != nil {
		return nil, erro
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API do Google retornou status: %d", resp.StatusCode)
	}

	corpo, erro := io.ReadAll(resp.Body)
	if erro != nil {
		return nil, erro
	}

	var dadosGoogle RespostaDoGoogle
	if erro := json.Unmarshal(corpo, &dadosGoogle); erro != nil {
		return nil, erro
	}

	var lista []interface{}
	for _, item := range dadosGoogle.Items {
		lista = append(lista, item.VolumeInfo)
	}

	return lista, nil
}
