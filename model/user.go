package model

import (
	"Api-Aula1/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	CPF      string `json:"cpf,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}
	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o e-mail inserido é inválido")
	}
	if err := utils.CPFValidator(u.CPF); err != nil {
		return err
	}
	if step == "create" && u.Password == "" {
		return errors.New("a senha é obrigatória")
	}
	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	u.Name = strings.ToLower(u.Name)

	return nil
}
