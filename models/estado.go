package models

import (
	"errors"

	"gorm.io/gorm"
)

type Estado struct {
	gorm.Model
	Status   string `json:"status,omitempty"`
	Nome string `json:"nome,omitempty"`
}



func (estado *Estado) Preparar() error {
	// Chama a função 
	err := estado.ValidarEstado()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (estado *Estado) ValidarEstado() error {
    // Valida o status
    if estado.Status != "ativo" && estado.Status != "inativo" {
        return errors.New("status inválido")
    }
    
    // Valida o nome
    if len(estado.Nome) < 3 {
        return errors.New("nome deve ter pelo menos 3 caracteres")
    }
    
    return nil
}