package repository

import (
	"Api-Aula1/model"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}
func (repo UsersRepo) Create(user model.User) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"INSERT INTO users (name, email, cpf, password) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(user.Name, user.Email, user.CPF, user.Password)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repo UsersRepo) FindAll() ([]model.User, error) {
	linhas, erro := repo.db.Query("SELECT id, name, email, cpf FROM users")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []model.User

	for linhas.Next() {
		var usuario model.User

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Email,
			&usuario.CPF,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
func (repo UsersRepo) Update(ID uint64, user model.User) error {
	statement, erro := repo.db.Prepare(
		"UPDATE users SET name = ?, email = ?, cpf = ? WHERE id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Email, user.CPF, ID); erro != nil {
		return erro
	}

	return nil
}

func (repo UsersRepo) Delete(ID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
