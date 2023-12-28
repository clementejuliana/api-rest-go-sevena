package models

import (
	"errors"
	"regexp"
)

type Login struct {
	Email        string `json:"email"`
	Senha        string `json:"senha"`
}

func (login *Login) Preparar() error {
	// Chama a função
	err := login.ValidateLogin()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (login *Login) ValidateLogin() error {
	// Valida se os campos obrigatórios estão preenchidos
	if login.Email == "" {
		return errors.New("Usuário e-mail é obrigatório")
	}

	if login.Senha == "" {
		return errors.New("Senha é obrigatória")
	}

	// Valida se o e-mail é válido
	if !isEmailValid(login.Email) {
		return errors.New("E-mail inválido")
	}

	// Valida se a senha é válida
	if len(login.Senha) < 6 {
		return errors.New("Senha deve ter pelo menos 6 caracteres")
	}

	return nil
}

func isEmailValid(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	match := regexp.MustCompile(regex).MatchString
	return match(email)
}
