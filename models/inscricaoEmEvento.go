package models

import (
	"time"

	"gorm.io/gorm"
)


type InscricaoEmEvento struct {
	gorm.Model
	Status              string    `json:"status,omitempty"`
	Data                time.Time `json:"data,omitempty"`
	Hora                time.Time `json:"hora,omitempty"`
	Evento              Evento    
	EventoID            uint       `json:"evento_id,omitempty"`
	Usuario             Usuario   
	UsuarioID           uint       `json:"usuario_id,omitempty"`

}