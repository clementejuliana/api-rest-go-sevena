package models

import (
	"errors"
	"time"
	"gorm.io/gorm"
)

type InscricaoEmAtividade struct {
	gorm.Model
	AtividadeID        int               `json:"atividade_id,omitempty"`
	EventoID           int               `json:"evento_id,omitempty"`
	Status             string            `json:"status,omitempty"`
	Data               time.Time         `json:"data,omitempty"`
	Hora               time.Time         `json:"hora,omitempty"`
	ControlePresencaID uint              `json:"controle_presenca_id,omitempty" gorm:"foreignKey:ControlePresenca"`
	ControlePresenca   *ControlePresenca `json:"controle_presenca,omitempty"`
	UsuarioID          int               `json:"usuario_id,omitempty"`
	Usuario            *Usuario          `json:"usuario,omitempty"`
	Atividade          *Atividade        `json:"atividade,omitempty"`
}

func (inscricaoA *InscricaoEmAtividade) Preparar(db *gorm.DB) error {
	// Chama a função ValidarInscricaoAtividade
	err := inscricaoA.ValidarInscricaoAtividade()
	if err != nil {
		return err
	}

	//Verifica conflito de horário passando a instância do banco de dados
	err = inscricaoA.VerificarConflitoHorario(db)
	if err != nil {
		return err
	}

	// Retorna nil se não houver erros
	return nil
}

func (i *InscricaoEmAtividade) ValidarInscricaoAtividade() error {
	// Verifica se o ID da atividade é válido
	if i.AtividadeID == 0 {
		return errors.New("atividade é obrigatório")
	}

	// Verifica se o ID do evento é válido
	if i.EventoID == 0 {
		return errors.New("evento é obrigatório")
	}

	// Verifica se o status é válido
	if i.Status != "pendente" && i.Status != "confirmada" && i.Status != "cancelada" {
		return errors.New("status deve ser 'pendente', 'confirmada' ou 'cancelada'")
	}

	// Verifica se a data é válida
	if i.Data.IsZero() {
		return errors.New("data é obrigatória")
	}

	// Verifica se a hora é válida
	if i.Hora.IsZero() {
		return errors.New("hora é obrigatória")
	}

	// Verifica se o controle de presença é válido
	if i.ControlePresencaID == 0 {
		return errors.New("controle presenca é obrigatório")
	}

	// Se chegamos até aqui, todos os campos são válidos
	return nil
}

func (i *InscricaoEmAtividade) VerificarConflitoHorario(db *gorm.DB) error {
	// Consulta para verificar conflito de horário
	var count int64
	result := db.Model(&InscricaoEmAtividade{}).
		Where("Evento_ID = ? AND Data = ? AND ((Hora >= ? AND Hora < ?) OR (Hora <= ? AND ? < Hora))",
			i.EventoID, i.Data, i.Hora, i.Hora.Add(time.Hour), i.Hora, i.Hora.Add(time.Hour)).
		Count(&count)

	if result.Error != nil {
		return result.Error
	}

	if count > 0 {
		return errors.New("conflito de horário com outra inscrição")
	}

	return nil
}
