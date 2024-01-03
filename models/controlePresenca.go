package models

import (
	"errors"

	"time"

	"gorm.io/gorm"
)

type ControlePresenca struct {
	gorm.Model
	Status        string    `json:"status,omitempty"`
	HoraEntrada   time.Time `json:"hora_entrada,omitempty"`
	HoraSaida     time.Time `json:"hora_saida,omitempty"`
	NotificacaoID uint      `json:"notificacao_id,omitempty"`
}

func (controle *ControlePresenca) Preparar() error {
	// Chama a função
	err := controle.ValidarControle()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (c *ControlePresenca) ValidarControle() error {
	if c.Status != "ativo" && c.Status != "inativo" {
		return errors.New("status inválido")
	}

	if c.HoraEntrada.IsZero() {
		return errors.New("hora entrada é obrigatória")
	}
	if c.HoraSaida.IsZero() {
		return errors.New("hora saida é obrigatória")
	}
	// adicione mais validações conforme necessário
	return nil
}

