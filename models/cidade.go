package models

type Cidade struct {
	CidadeID int `json:"cidade_id,omitempty"`
	Status   string `json:"status,omitempty"`
	Nome string`json:"nome,omitempty"`
	EstadoID int `json:"estado_id,omitempty"`

}