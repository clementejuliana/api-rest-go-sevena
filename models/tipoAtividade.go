package models

type TipoAtividade struct {
	TipoAtividadeID int `json:"tipo_atividade_id,omitempty"`
	Status string `json:"status,omitempty"`
	TipoDaAtividade string `json:"tipodaAtividade,omitempty"`
}