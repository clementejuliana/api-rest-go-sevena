package models

import (
	"time"

	"gorm.io/gorm"
)

type Atividade struct {
	gorm.Model
    AtividadeID        int       `json:"atividade_id,omitempty"`
    Status             string    `json:"status,omitempty"`
    TipoAtividade      TipoAtividade `gorm:"foreignKey:TipoAtividadeID,omitempty"`
    TipoAtividadeID    int       `json:"tipo_atividade_id,omitempty"`
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
    QuantidadeInscritos int      `json:"quantidade_inscritos,omitempty"`
    Local              Local     `gorm:"foreignKey:LocalID,omitempty"`
    LocalID            int       `json:"local_id,omitempty"`
}
