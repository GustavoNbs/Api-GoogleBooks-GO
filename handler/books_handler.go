package handler

import (
	"encoding/json"
	"fmt"
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

	// Pega o parâmetro "nome" da URL
	nomePraPesquisar := r.URL.Query().Get("nome")

	// Valida se o parâmetro foi enviado
	if nomePraPesquisar == "" {
		http.Error(w, "Ei, você esqueceu de mandar o parametro 'nome' na URL!", 400)
		return
	}

	// Chama a função que busca no Google
	livros, erro := buscarNoGoogle(nomePraPesquisar)
	if erro != nil {
		http.Error(w, "Deu erro ao falar com o Google: "+erro.Error(), 500)
		return
	}

	// Retorna o resultado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)

}

func buscarNoGoogle(nome string) ([]interface{}, error) {
	// Arruma o nome pra não quebrar a URL (espaços viram %20)
	nomeFormatado := url.QueryEscape(nome)
	urlGoogle := "https://www.googleapis.com/books/v1/volumes?q=" + nomeFormatado

	resp, erro := http.Get(urlGoogle)
	if erro != nil {
		return nil, erro
	}
	defer resp.Body.Close()

	var dadosGoogle RespostaDoGoogle
	if erro := json.NewDecoder(resp.Body).Decode(&dadosGoogle); erro != nil {
		return nil, erro
	}

	var lista []interface{}
	for _, item := range dadosGoogle.Items {
		lista = append(lista, item.VolumeInfo)
	}

	return lista, nil
}
