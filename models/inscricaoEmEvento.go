package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type InscricaoEmEvento struct {
	gorm.Model
	Status    string    `json:"status,omitempty"`
	Data      time.Time `json:"data,omitempty"`
	Hora      time.Time `json:"hora,omitempty"`
	Evento    Evento
	EventoID  uint `json:"evento_id,omitempty"`
	Usuario   Usuario
	UsuarioID uint `json:"usuario_id,omitempty"`
}

func (inscricaoEvento *InscricaoEmEvento) Preparar() error {
	// Chama a função
	err := inscricaoEvento.ValidarInscricaoEvento()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (i *InscricaoEmEvento) ValidarInscricaoEvento() error {
	// Verifica se o status é válido
	if i.Status != "ativo" && i.Status != "inativo" {
		return errors.New("status é obrigatório e deve ser 'ativo' ou 'inativo'")
	}

	// Verifica se a data é válida
	if i.Data.IsZero() {
		return errors.New("data é obrigatória")
	}

	// Verifica se a hora é válida
	if i.Hora.IsZero() {
		return errors.New("hora é obrigatória")
	}

	// Verifica se o ID do evento é válido
	if i.EventoID == 0 {
		return errors.New("evento é obrigatório")
	}
	

	// Verifica se o ID do usuário é válido
	if i.UsuarioID == 0 {
		return errors.New("usuario é obrigatório")
	}

	return nil
}
