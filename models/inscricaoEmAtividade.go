package models

import (
	"time"

	"gorm.io/gorm"
)

type InscricaoEmAtividade struct {
	gorm.Model
    AtividadeID        int       `json:"atividade_id,omitempty"`
	EventoID           int       `json:"evento_id,omitempty"`
	Status             string    `json:"status,omitempty"`
	Data               time.Time `json:"data,omitempty"`
	Hora               time.Time `json:"hora,omitempty"`
	ControlePresenca   ControlePresenca  
	ControlePresencaID int       `json:"controle_presenca_id,omitempty"`
	
}
