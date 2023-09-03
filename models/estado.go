package models

type Estado struct {
	Status   string `json:"status,omitempty"`
	EstadoID int`json:"estado_id,omitempty"`
	Nome string `json:"nome,omitempty"`

}