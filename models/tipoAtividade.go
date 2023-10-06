package models

import "gorm.io/gorm"

type TipoAtividade struct {
	gorm.Model
    TipoAtividadeID int       `json:"tipo_atividade_id,omitempty"`
    Status          string    `json:"status,omitempty"`
    TipoDaAtividade string    `json:"tipo_da_atividade,omitempty"`
}