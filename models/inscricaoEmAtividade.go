package models

import "time"

type InscricaoEmAtividade struct {
	AtividadeID int `json:"atividade_id,omitempty"`
	EventoID    int `json:"evento_id,omitempty"`
	Status string `json:"status,omitempty"`
	Data time.Time `json:"data,omitempty"`
	Hora time.Time `json:"hora,omitempty"`
	ControlePresencaID int `json:"contole_presenca_id,omitempty"`
	
}
