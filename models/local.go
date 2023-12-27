package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Local struct {
	gorm.Model
	Status      string      `json:"status,omitempty"`
	Sala        string      `json:"sala,omitempty"`
	Setor       string      `json:"setor,omitempty"`
	Atividades  []Atividade `json:"atividades" gorm:"foreignKey:LocalID"`
	Eventos     []Evento    `json:"eventos" gorm:"foreignKey:LocalID"`
	DataHoraFim time.Time   `json:"dataHoraFim,omitempty"`
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
	if local.Status == "Disponível" {
		// Verifica se a sala já está reservada para outro evento
		for _, evento := range local.Eventos {
			if evento.HoraInicio.Before(local.DataHoraFim) && evento.HoraFim.After(local.DataHoraFim) {
				return errors.New("Sala já está reservada para outro evento no mesmo dia e horário")
			}
		}
	}
	if local.Sala == "" {
		return errors.New("Sala é obrigatória")
	}

	if local.Setor == "" {
		return errors.New("Setor é obrigatório")
	}

	return nil
}

// // Função para obter horários disponíveis em um determinado dia
// func (local *Local) HorariosDisponiveis(data time.Time) ([]time.Time, error) {
// 	// Lógica para obter os horários disponíveis na data especificada
// 	// Pode envolver consulta ao banco de dados, por exemplo

// 	// Exemplo: retornando uma lista fixa de horários para fins de demonstração
// 	horarios := []time.Time{
// 		data.Add(1 * time.Hour),
// 		data.Add(2 * time.Hour),
// 		data.Add(3 * time.Hour),
// 		// Adicione mais horários conforme necessário
// 	}

// 	return horarios, nil
// }
