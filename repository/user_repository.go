package repository

import (
	"Api-Aula1/model"
	"errors"
	"fmt"
)

var users []model.User
var nextID = 1

func Create(user model.User) model.User {
	user.ID = fmt.Sprintf("%d", nextID)
	nextID++
	users = append(users, user)
	return user
}

func FindAll() []model.User {
	return users
}

func Update(id string, updatedUser model.User) (model.User, error) {
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = user.ID
			users[i] = updatedUser
			return updatedUser, nil
		}
	}
	return model.User{}, errors.New("utilizador não encontrado")
}

func Delete(id string) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("utilizador não encontrado")
}
