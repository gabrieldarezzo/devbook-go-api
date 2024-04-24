package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represent socialNetwork
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will calling methods validate and format to santize User
func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

// TODO: Refactor with Strategy
func (user *User) validate(stage string) error {

	if stage == "NEW_USER" && user.Password == "" {
		return errors.New("O campo: 'password' é obrigatório e não pode estar em branco")
	}

	if user.Name == "" {
		return errors.New("O campo: 'name' é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O campo: 'nick' é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O campo: 'email' é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O campo: 'email' é invalido")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "NEW_USER" {
		passwordHashed, erro := security.GenerateHash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordHashed)
	}
	return nil
}
