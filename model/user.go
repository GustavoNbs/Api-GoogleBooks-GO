package model

import (
	"strconv"
	"strings"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func (u *User) ValidateCPF() bool {
	cpf := strings.ReplaceAll(strings.ReplaceAll(u.CPF, ".", ""), "-", "")

	if len(cpf) != 11 {
		return false
	}

	allEqual := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (10 - i)
	}
	remainder := sum % 11
	digit1 := 0
	if remainder >= 2 {
		digit1 = 11 - remainder
	}

	currentDigit, _ := strconv.Atoi(string(cpf[9]))
	if currentDigit != digit1 {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (11 - i)
	}
	remainder = sum % 11
	digit2 := 0
	if remainder >= 2 {
		digit2 = 11 - remainder
	}

	currentDigit, _ = strconv.Atoi(string(cpf[10]))
	return currentDigit == digit2
}
