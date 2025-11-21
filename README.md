# API de Busca de Livros (Go + Google Books)

Este projeto é uma API desenvolvida em **Go (Golang)** que permite pesquisar livros diretamente na API do Google Books através de um termo enviado pelo usuário.

O sistema funciona como um "proxy": ele recebe sua requisição, valida o parâmetro de busca, consulta o Google Books e retorna os resultados formatados para JSON.

---

## Integrantes do Grupo

* **Nome do Aluno 1:** [Gustavo Biscoto Nebes]
* **Nome do Aluno 2:** [Davi Almeida Resende]
* **Nome do Aluno 3:** [Leonardo Sena Machado Durães]


---

## Tecnologias e APIs

* **Linguagem:** Go (Golang)
* **Roteamento:** Gorilla Mux
* **API Externa:** [Google Books API](https://developers.google.com/books)

---

## Como Rodar o Projeto

### Pré-requisitos

* Ter a linguagem [Go (Golang)](https://go.dev/dl/) instalada na máquina.

### Passo a Passo

1. **Abra o terminal** na pasta do projeto.
2. **Baixe as dependências** (caso necessário):

   ```bash
   ```

go mod tidy

````
3. **Inicie o servidor:**
```bash
go run main.go
````

Você verá a mensagem no terminal: `Servidor ouvindo em :8080 ...`

---

##  Como Testar

A rota principal da aplicação é: `GET /books/search`

 **Importante:** Você deve obrigatoriamente passar o parâmetro **nome** na URL para que a busca aconteça.

### 1. Pelo Navegador

Basta acessar o link abaixo (você pode alterar o nome do livro no final):

```
http://localhost:8080/books/search?nome=Dom Quixote
```

### 2. Pelo Postman / Insomnia

1. Crie uma requisição do tipo **GET**.
2. Insira a URL:

```
http://localhost:8080/books/search
```

3. Adicione o Parâmetro de consulta (Query Param):

   * **Chave (Key):** nome
   * **Valor (Value):** Clean Code (ou o livro que desejar).
4. Clique em **Send**.

---

## Exemplos de Resposta

### Sucesso (200 OK):

```json
[
    {
        "title": "Harry Potter e a Ordem da Fénix",
        "authors": ["J. K. Rowling"],
        "description": "Este tem sido um Verão ainda mais insuportavel que o costume...",
        "pageCount": 750
    }
]
```

### Erro - Falta de parâmetro (400 Bad Request):

```
Ei, você esqueceu de mandar o parametro 'nome' na URL!
```

---

## Configuração de Porta

O servidor roda por padrão na porta **:8080**. Se precisar alterar:

1. Abra o arquivo `main.go`.
2. Localize a constante `addr`.
3. Altere o valor (ex: `":"3000"`).
4. Reinicie o servidor.
