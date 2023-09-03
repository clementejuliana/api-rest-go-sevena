package models

import "time"

type Atividade struct {
	AtividadeID        int       `json:"atividade_id,omitempty"`
	Status             string    `json:"status,omitempty"`
	TipoAtividadeID    int       `json:"tipoAtividade_id,omitempty"`
	Titulo             string    `json:"titulo,omitempty"`
	Resumo             string    `json:"resumo,omitempty"`
	Data               time.Time `json:"data,omitempty"`
	HoraInicio         time.Time `json:"hora_inicio,omitempty"`
	HoraTermino        time.Time `json:"hora_termino,omitempty"`
	ValorInscricao     float64   `json:"valor_inscricao"`
	Observacao         string    `json:"observacao"`
	Ministrante        string    `json:"ministrante,omitempty"`
	QuantidadeVagas    int       `json:"quantidade_vagas,omitempty"`
	Duracao            float64   `json:"duracao,omitempty"`
	CargaHoraria       int       `json:"carga_horaria,omitempty"`
	QuantidadeInscrios int       `json:"quantidade_inscritos,omitempty"`
	LocalID            int       `json:"local_id,omitempty"`
}
