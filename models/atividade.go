package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Atividade struct {
	gorm.Model
	Status              string  `json:"status,omitempty"`
	TipoAtividade       string  `json:"tipo_atividade,omitempty"`
	Titulo              string  `json:"titulo,omitempty"`
	Resumo              string  `json:"resumo,omitempty"`
	Data                string  `json:"data,omitempty"`
	HoraInicio          string  `json:"hora_inicio,omitempty"`
	HoraTermino         string  `json:"hora_termino,omitempty"`
	ValorInscricao      float64 `json:"valor_inscricao"`
	Observacao          string  `json:"observacao"`
	Ministrante         string  `json:"ministrante,omitempty"`
	QuantidadeVagas     int     `json:"quantidade_vagas,omitempty"`
	Duracao             float64 `json:"duracao,omitempty"`
	CargaHoraria        int     `json:"carga_horaria,omitempty"`
	QuantidadeInscritos int     `json:"quantidade_inscritos,omitempty"`
	LocalID             int     `json:"local_id" gorm:"foreignKey:LocalID"`
	Local               Local   `json:"local,omitempty"`
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
	

	if a.Status != "ativo" && a.Status != "inativo" {
		return errors.New("status é obrigatório")
	}

	if a.Titulo == "" {
		return errors.New("titulo é obrigatório")
	}
	if a.Resumo == "" {
		return errors.New("resumo é obrigatório")
	}



	// // Validar a hora de início
	// horaInicio, err := time.Parse("15:04", a.HoraInicio)
	// if err != nil || horaInicio.Before(now) {
	// 	return errors.New("hora_inicio é obrigatória e deve estar no futuro")
	// }

	// // Validar a hora de término
	// horaTermino, err := time.Parse("15:04", a.HoraTermino)
	// if err != nil || horaTermino.Before(horaInicio) {
	// 	return errors.New("hora_termino é obrigatória e deve ser após hora_inicio")
	// }

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
	// Converte as strings de hora para objetos time.Time
	horaInicio, err := time.Parse("15:04", a.HoraInicio)
	if err != nil {
		return Intervalo{}, errors.New("hora_inicio inválida")
	}

	horaTermino, err := time.Parse("15:04", a.HoraTermino)
	if err != nil {
		return Intervalo{}, errors.New("hora_termino inválida")
	}

	if horaInicio.After(horaTermino) {
		return Intervalo{}, errors.New("hora_inicio deve ser antes de hora_termino")
	}

	return Intervalo{
		DataInicio: horaInicio,
		DataFim:    horaTermino,
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
