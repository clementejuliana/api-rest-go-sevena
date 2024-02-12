package models

import (
	"errors"

	"gorm.io/gorm"
)

type Local struct {
	gorm.Model
	Status      string `json:"status,omitempty"`
	Sala        string `json:"sala,omitempty"`
	Setor       string `json:"setor,omitempty"`
	DataHoraFim string `json:"dataHoraFim,omitempty"`
}

func (local *Local) Preparar() error {
	// Chama a função ValidarU
	err := local.ValidateLocal()
	// Verifica se houve erros
	if err != nil {
		return err
	}

	// Retorna nil se não houver erros
	return nil
}

func (local *Local) ValidateLocal() error {
	if local.Status != "ativo" && local.Status != "inativo" {
		return errors.New("status é obrigatório")
	}

	if local.Status != "ativo" && local.Status != "inativo" {
		return errors.New("status é obrigatório")
	}
	if local.Sala == "" {
		return errors.New("Sala é obrigatória")
	}

	if local.Setor == "" {
		return errors.New("Setor é obrigatório")
	}

	return nil
}
