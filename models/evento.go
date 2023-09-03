package models

import "time"

type Evento struct {
	EventoID int `json:"evento_id,omitempty"`
	Status string `json:"status,omitempty"`
	Nome string `json:"nome,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	DataFinal  time.Time`json:"data_final,omitempty"`
	LocalID int `json:"local_id,omitempty"`
}