package models

import (
	"time"

	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
    EventoID  int       `json:"evento_id,omitempty"`
    Status    string    `json:"status,omitempty"`
    Nome      string    `json:"nome,omitempty"`
    Descricao string    `json:"descricao,omitempty"`
    DataInicio time.Time `json:"data_inicio,omitempty"`
    DataFinal time.Time `json:"data_final,omitempty"`
	LocalID   int       `json:"local_id" gorm:"foreignKey:LocalID"`
    Local     Local     `json:"local,omitempty"`
}