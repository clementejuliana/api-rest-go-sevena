package models

import (
	"errors"

	"gorm.io/gorm"
)

type TipoAtividade struct {
    gorm.Model
    Nome string `json:"nome"`
    Status      string    `json:"status,omitempty"`
}

func (tipoAtividade *TipoAtividade) Preparar() error {
	// Chama a função
	err := tipoAtividade.ValidateTipoAtividade()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}


func (tipoAtividade *TipoAtividade) ValidateTipoAtividade() error {
    // Valida se os campos obrigatórios estão preenchidos
    if tipoAtividade.Nome == "" {
        return errors.New("Nome é obrigatório")
    }

    // Valida se o nome é válido
    if len(tipoAtividade.Nome) < 3 {
        return errors.New("Nome deve ter pelo menos 3 caracteres")
    }

    return nil
}
