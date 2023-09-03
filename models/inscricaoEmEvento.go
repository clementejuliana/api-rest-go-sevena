package models

import "time"


type InscricaoEmEvento struct {
	InscricaoEmEventoID int `json:"inscricao_em_evento_id,omitempty"`
	Status string `json:"status,omitempty"`
	Data time.Time `json:"data,omitempty"`
	Hora time.Time `json:"hora,omitempty"`
	EventoID int `json:"evento_id,omitempty"`
	UsuarioID int `json:"usuario_id,omitempty"`

}