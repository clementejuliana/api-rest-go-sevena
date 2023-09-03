package models

import "time"

type ControlePresenca struct {
	ControlePresencaID int `json:"contole_presenca_id,omitempty"`
	Status string `json:"status,omitempty"`
	HoraEntrada time.Time `json:"hora_entrada,omitempty"`
	HoraSaida time.Time `json:"hora_saida,omitempty"`
}