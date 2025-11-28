# API de Busca de Livros e Gestão de Usuários (Go + Google Books)

Este projeto é uma API desenvolvida em **Go (Golang)** que possui duas funcionalidades principais:

1. Pesquisar livros diretamente na API do Google Books.  
2. Gerir um cadastro de usuários (CRUD) com validação de CPF.

O sistema funciona como um "proxy" para livros e possui um banco de dados em memória para os usuários.

---

## Integrantes do Grupo

* **Nome do Aluno 1:** Gustavo Biscoto Nebes  
* **Nome do Aluno 2:** Davi Almeida Resende  
* **Nome do Aluno 3:** Leonardo Sena Machado Durães

---

## Tecnologias e APIs

| Componente       | Utilizado |
|------------------|-----------|
| Linguagem        | Go (Golang) |
| Roteamento       | Gorilla Mux |
| API Externa      | Google Books API |
| Arquitetura      | MVC (Model, View/Handler, Controller, Repository) |

---

## Como Rodar o Projeto

### Pré-requisitos

* Ter a linguagem **Go (Golang)** instalada na máquina.  
  Download: https://go.dev/dl/

---

### Passo a Passo

1. **Abra o terminal** na pasta do projeto.
2. **Baixe as dependências** (caso necessário):

   ```bash
   go mod tidy
   ```

3. **Inicie o servidor:**

   ```bash
   go run main.go
   ```

Você verá a mensagem:  
**Servidor ouvindo em :8080 ...**

---

## Como Testar – Busca de Livros

**Rota:** `GET /books/search`  
É obrigatório passar o parâmetro `nome`.

**Exemplo no navegador:**

```
http://localhost:8080/books/search?nome=Dom Quixote
```

---

## Como Testar o CRUD de Usuários

A API permite **Criar, Listar, Atualizar e Deletar usuários**, com validação de CPF matematicamente correta.

---

### 1. Criar Usuário – POST

**URL:**

```
http://localhost:8080/users
```

**Body (JSON):**

```json
{
    "name": "Maria Silva",
    "email": "maria@email.com",
    "cpf": "52998224725"
}
```

*Use um CPF válido — CPFs inválidos retornam **400 Bad Request***.

---

### 2. Listar Usuários – GET

**URL:**

```
http://localhost:8080/users
```

**Resposta esperada:**

```json
[
    {
        "id": "1",
        "name": "Maria Silva",
        "email": "maria@email.com",
        "cpf": "52998224725"
    }
]
```

---

### 3. Atualizar Usuário – PUT

Substitua `{id}` (ex: 1)  

```
http://localhost:8080/users/{id}
```

**Body (JSON):**

```json
{
    "name": "Maria Silva Atualizada",
    "email": "maria.novo@email.com",
    "cpf": "52998224725"
}
```

---

### 4. Deletar Usuário – DELETE

```
http://localhost:8080/users/{id}
```

**Resposta:** Status `204 No Content`

---

## Códigos de Erro Comuns

| Código | Significado |
|--------|-------------|
| **400 Bad Request** | Parâmetros faltando ou CPF inválido |
| **404 Not Found** | Usuário não encontrado ao atualizar/deletar |
| **500 Internal Server Error** | Falha ao conectar na API externa |

---

## Configuração de Porta

O servidor roda em `:8080` por padrão. Para alterar:

1. Abra `main.go`  
2. Localize a constante `addr`  
3. Altere o valor (ex: `":3000"`)  
4. Reinicie o servidor
