package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ControlePresenca struct {
	gorm.Model
	Status               string                `json:"status,omitempty"`
	InscricaoAtividade   uint                  `json:"inscricao_atividade"`
	HoraEntrada          string                `json:"hora_entrada,omitempty"`
	HoraSaida            string                `json:"hora_saida,omitempty"`
	NotificacaoID        uint                  `json:"notificacao_id,omitempty"`
	InscricaoEmAtividade *InscricaoEmAtividade `json:"inscricao,omitempty" `
}

//gorm:"foreignKey:InscricaoAtividade"
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
	if c.HoraEntrada == "" {
		return errors.New("hora entrada é obrigatória")
	}
	if c.HoraSaida == "" {
		return errors.New("hora saida é obrigatória")
	}

	// Validar as datas
	_, err := time.Parse("2006-01-02T15:04:05", c.HoraEntrada)
	if err != nil {
		return errors.New("hora entrada inválida")
	}

	_, err = time.Parse("2006-01-02T15:04:05", c.HoraSaida)
	if err != nil {
		return errors.New("hora saida inválida")
	}
	// adicione mais validações conforme necessário
	return nil
}
