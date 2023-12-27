package models

import (
	"errors"
	
	"time"
)

type Login struct {
    ID        int    `json:"id"`
    UsuarioEmail string `json:"usuario_email"`
    Status    string `json:"status"`
    ExpirationData time.Time `json:"expiration_data"`
    Email string `json:"email"`
    Senha string `json:"senha"`
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
    if login.UsuarioEmail == "" {
        return errors.New("Usuário e-mail é obrigatório")
    }

    if login.Senha == "" {
        return errors.New("Senha é obrigatória")
    }

    // Valida se o e-mail é válido
    //if !strings.Contains(validEmails, login.Email) {
        //return errors.New("E-mail inválido")
   // }

    // Valida se a senha é válida
    if len(login.Senha) < 6 {
        return errors.New("Senha deve ter pelo menos 6 caracteres")
    }

    return nil
}
