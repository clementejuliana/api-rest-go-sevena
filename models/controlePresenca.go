package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlePresenca struct {
	gorm.Model
    ControlePresencaID int       `json:"controle_presenca_id,omitempty"`
    Status             string    `json:"status,omitempty"`
    HoraEntrada        time.Time `json:"hora_entrada,omitempty"`
    HoraSaida          time.Time `json:"hora_saida,omitempty"`
}