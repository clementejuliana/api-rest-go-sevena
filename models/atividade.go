package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Atividade struct {
	gorm.Model         
	Status              string        `json:"status,omitempty"`
	TipoAtividade       TipoAtividade `json:"tipo_atividade,omitempty"`
	TipoAtividadeID     int           `json:"tipo_atividade_id,omitempty"`
	Titulo              string        `json:"titulo,omitempty"`
	Resumo              string        `json:"resumo,omitempty"`
	Data                time.Time     `json:"data,omitempty"`
	HoraInicio          time.Time     `json:"hora_inicio,omitempty"`
	HoraTermino         time.Time     `json:"hora_termino,omitempty"`
	ValorInscricao      float64       `json:"valor_inscricao"`
	Observacao          string        `json:"observacao"`
	Ministrante         string        `json:"ministrante,omitempty"`
	QuantidadeVagas     int           `json:"quantidade_vagas,omitempty"`
	Duracao             float64       `json:"duracao,omitempty"`
	CargaHoraria        int           `json:"carga_horaria,omitempty"`
	QuantidadeInscritos int           `json:"quantidade_inscritos,omitempty"`
	LocalID             int           `json:"local_id" gorm:"foreignKey:LocalID"`
	Local               Local         `json:"local,omitempty"`
}

func (atividade *Atividade) Preparar() error {
	// Chama a função ValidarU
	err := atividade.ValidarAtividade()
	// Verifica se houve erros
	if err != nil {
		return err
	}

	// Retorna nil se não houver erros
	return nil
}

type Intervalo struct {
	DataInicio time.Time
	DataFim    time.Time
}

func (a *Atividade) ValidarAtividade() error {
	now := time.Now()

	if a.Status != "ativo" && a.Status != "inativo" {
		return errors.New("status é obrigatório")
	}
	if a.TipoAtividadeID == 0 {
		return errors.New("tipo_atividade_id é obrigatório")
	}
	if a.Titulo == "" {
		return errors.New("titulo é obrigatório")
	}
	if a.Resumo == "" {
		return errors.New("resumo é obrigatório")
	}
	if a.Data.IsZero() || a.Data.Before(now) {
		return errors.New("data é obrigatória e deve estar no futuro")
	}
	if a.HoraInicio.IsZero() || a.HoraInicio.Before(now) {
		return errors.New("hora_inicio é obrigatória e deve estar no futuro")
	}
	if a.HoraTermino.IsZero() || a.HoraTermino.Before(a.HoraInicio) {
		return errors.New("hora_termino é obrigatória e deve ser após hora_inicio")
	}
	if a.ValorInscricao == 0 {
		return errors.New("valor_inscricao é obrigatório")
	}
	if a.Ministrante == "" {
		return errors.New("ministrante é obrigatório")
	}
	if a.QuantidadeVagas == 0 {
		return errors.New("quantidade_vagas é obrigatório")
	}
	if a.QuantidadeInscritos > a.QuantidadeVagas {
		return errors.New("a quantidade de inscritos não pode ser maior que a quantidade de vagas disponíveis")
	}
	if a.Duracao == 0 {
		return errors.New("duracao é obrigatório")
	}
	if a.CargaHoraria == 0 {
		return errors.New("carga_horaria é obrigatório")
	}
	if a.LocalID == 0 {
		return errors.New("local_id é obrigatório")
	}
	return nil
}

func (a *Atividade) GetIntervalo() (Intervalo, error) {
	if a.HoraInicio.After(a.HoraTermino) {
		return Intervalo{}, errors.New("hora_inicio deve ser antes de hora_termino")
	}
	return Intervalo{
		DataInicio: a.HoraInicio,
		DataFim:    a.HoraTermino,
	}, nil
}

func (a *Atividade) Inscricao() error {
	if a.QuantidadeInscritos >= a.QuantidadeVagas {
		return errors.New("Não há vagas disponíveis para esta atividade")
	}
	a.QuantidadeInscritos++
	// Salve a atividade atualizada no banco de dados
	return nil
}
