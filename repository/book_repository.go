package repository

import (
	"Api-Aula1/model"
	"database/sql"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db}
}

func (repo BooksRepo) Salvar(book model.Book) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"INSERT INTO books (title, authors, description, user_id) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(book.Title, book.Authors, book.Description, book.UserID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repo BooksRepo) Buscar(userID uint64) ([]model.Book, error) {
	linhas, erro := repo.db.Query("SELECT id, title, authors, description FROM books WHERE user_id = ?", userID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var books []model.Book
	for linhas.Next() {
		var book model.Book
		if erro = linhas.Scan(&book.ID, &book.Title, &book.Authors, &book.Description); erro != nil {
			return nil, erro
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo BooksRepo) Atualizar(bookID uint64, userID uint64, book model.Book) error {
	statement, erro := repo.db.Prepare(
		"UPDATE books SET title = ?, authors = ?, description = ? WHERE id = ? AND user_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(book.Title, book.Authors, book.Description, bookID, userID); erro != nil {
		return erro
	}
	return nil
}

func (repo BooksRepo) Deletar(bookID uint64, userID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM books WHERE id = ? AND user_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(bookID, userID); erro != nil {
		return erro
	}
	return nil
}
