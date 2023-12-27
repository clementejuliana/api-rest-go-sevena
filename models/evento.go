package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	//  EventoID  int       `json:"evento_id,omitempty"`
	Status     string    `json:"status,omitempty"`
	Nome       string    `json:"nome,omitempty"`
	Descricao  string    `json:"descricao,omitempty"`
	DataInicio time.Time `json:"data_inicio,omitempty"`
	DataFinal  time.Time `json:"data_final,omitempty"`
	LocalID    int       `json:"local_id" gorm:"foreignKey:LocalID"`
	Local      Local     `json:"local,omitempty"`
	HoraInicio time.Time `json:"horaInicio"`
	HoraFim    time.Time `json:"horaFim"`
}

func (evento *Evento) Preparar(db *gorm.DB) error {
	// Chama a função ValidarEvento
	err := evento.ValidarEvento()
	if err != nil {
		return err
	}
//Verifica conflito de horário passando a instância do banco de dados
err = evento.VerificarConflitosHorario(db)
if err != nil {
	return err
}
	
	return nil
}

func (e *Evento) ValidarEvento() error {

	if e.Status != "ativo" && e.Status != "inativo" {
		return errors.New("status deve ser 'ativo' ou 'inativo'")
	}
	if e.Nome == "" {
		return errors.New("nome é obrigatório")
	}
	if len(e.Nome) < 5 {
		return errors.New("nome deve ter pelo menos 5 caracteres")
	}
	if len(e.Nome) > 255 {
		return errors.New("nome deve ter no máximo 255 caracteres")
	}
	if e.Descricao == "" {
		return errors.New("descrição é obrigatória")
	}
	if len(e.Descricao) < 100 {
		return errors.New("descrição deve ter pelo menos 100 caracteres")
	}
	if len(e.Descricao) > 10000 {
		return errors.New("descrição deve ter no máximo 10.000 caracteres")
	}
	if e.DataInicio.After(e.DataFinal) {
		return errors.New("data de inicio deve ser antes de data final do evento")
	}
	if e.LocalID == 0 {
		return errors.New("local  é obrigatório")
	}
	return nil
}

func (e *Evento) VerificarConflitosHorario(db *gorm.DB) error {
	// Consulta o banco de dados para verificar se há conflitos de horário
	var count int64
	err := db.Model(&Evento{}).
		Where("local_id = ? AND ((data_inicio <= ? AND data_final >= ?) OR (data_inicio <= ? AND data_final >= ?))",
			e.LocalID, e.DataInicio, e.DataInicio, e.DataFinal, e.DataFinal).
		Not("ID = ?", e.ID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("conflito de horário encontrado. Escolha outro horário ou local.")
	}

	return nil
}

func (evento *Evento) HorariosDisponiveis(db *gorm.DB) (_[]Local, data_inicio error) {
	// Substitua com sua própria lógica para obter os horários disponíveis para o local e data do evento
	// Considere consultar eventos existentes para o local e calcular os horários disponíveis

	// Exemplo de consulta de eventos para o local e data do evento
	var eventos []Evento
	err := db.Where("LocalID = ? AND (DataInicio >= ? AND DataFinal <= ?)", evento.LocalID, evento.DataInicio, evento.DataFinal).
		Find(&eventos).Error
	if err != nil {
		return nil, err
	}

	// Implemente a lógica para calcular os horários disponíveis com base nos eventos existentes

	// Retorne os horários disponíveis
	return []HorarioDisponivel{
		{HoraInicio: time.Now(), HoraFim: time.Now().Add(time.Hour)},
		{HoraInicio: time.Now().Add(2 * time.Hour), HoraFim: time.Now().Add(3 * time.Hour)},
		// Adicione mais horários disponíveis conforme necessário
	}, nil
}