package models

import (
	"time"

	"gorm.io/gorm"
)


type InscricaoEmEvento struct {
	gorm.Model
	InscricaoEmEventoID int       `json:"inscricao_em_evento_id,omitempty"`
	Status              string    `json:"status,omitempty"`
	Data                time.Time `json:"data,omitempty"`
	Hora                time.Time `json:"hora,omitempty"`
	Evento              Evento     `gorm:"foreignKey:EventoID,omitempty"`
	EventoID            int       `json:"evento_id,omitempty"`
	Usuario             Usuario   `gorm:"foreignKey:UsuarioID,omitempty"`
	UsuarioID           int       `json:"usuario_id,omitempty"`

}