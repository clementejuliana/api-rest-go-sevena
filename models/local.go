package models

import "gorm.io/gorm"

type Local struct {
	gorm.Model
    Status  string    `json:"status,omitempty"`
    Sala    string    `json:"sala,omitempty"`
    Setor   string    `json:"setor,omitempty"`
    Atividades []Atividade `gorm:"foreignKey:LocalID"`
    Eventos []Evento `gorm:"foreignKey:LocalID"`
}
