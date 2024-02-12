package models

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type InscricaoEmAtividade struct {
	gorm.Model
	AtividadeID        int               `json:"atividade_id,omitempty"`
	EventoID           int               `json:"evento_id,omitempty"`
	Status             string            `json:"status,omitempty"`
	Data               string            `json:"data,omitempty"`
	Hora               string            `json:"hora,omitempty"`
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

	// Validação para aceitar apenas números na data
	if match, _ := regexp.MatchString("[0-9]{2}/[0-9]{2}/[0-9]{4}$", i.Data); !match {
		return errors.New("data deve conter apenas números")
	}

	// Validação para aceitar apenas números na hora
	if match, _ := regexp.MatchString("[0-9]{2}:[0-9]{2}$", i.Hora); !match {
		return errors.New("hora deve conter apenas números")
	}

	// Verifica se o controle de presença é válido
	if i.ControlePresencaID == 0 {
		return errors.New("controle presenca é obrigatório")
	}

	// Se chegamos até aqui, todos os campos são válidos
	return nil
}

func (inscricaoA *InscricaoEmAtividade) VerificarConflitoHorario(db *gorm.DB) error {
	var inscricoes []InscricaoEmAtividade

	// Busca todas as inscrições do usuário no mesmo dia
	db.Where("usuario_id = ? AND data = ?", inscricaoA.UsuarioID, inscricaoA.Data).Find(&inscricoes)

	// Converte a hora de início da inscrição atual para um objeto Time
	horaInicioAtual, err := time.Parse("15:04", inscricaoA.Hora)
	if err != nil {
		return errors.New("hora de início inválida")
	}

	// Calcula a hora de término da inscrição atual com base na hora de início e na duração da atividade (vamos assumir que a duração é de 1 hora)
	horaTerminoAtual := horaInicioAtual.Add(time.Hour)

	for _, inscricao := range inscricoes {
		// Converte a hora de início da inscrição existente para um objeto Time
		horaInicioExistente, err := time.Parse("15:04", inscricao.Hora)
		if err != nil {
			return errors.New("hora de início existente inválida")
		}

		// Calcula a hora de término da inscrição existente com base na hora de início e na duração da atividade (assumindo 1 hora)
		horaTerminoExistente := horaInicioExistente.Add(time.Hour)

		// Verifica se há sobreposição entre as atividades
		if (horaInicioAtual.Before(horaTerminoExistente) && horaTerminoAtual.After(horaInicioExistente)) ||
			(horaInicioExistente.Before(horaTerminoAtual) && horaTerminoExistente.After(horaInicioAtual)) {
			return errors.New("conflito de horário detectado")
		}
	}

	return nil
}
