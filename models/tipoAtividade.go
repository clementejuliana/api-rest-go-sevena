package models

import "gorm.io/gorm"

type TipoAtividade struct {
	gorm.Model
    Status          string    `json:"status,omitempty"`
    TipoDaAtividade string    `json:"tipo_da_atividade,omitempty"`
}