package models

import "gorm.io/gorm"

type Cidade struct {
	gorm.Model
	Status   string `json:"status,omitempty"`
	Nome string`json:"nome,omitempty"`
	InstituicaoID int
	EstadoID int `json:"estado_id,omitempty"`
	Usuarios []Usuario `gorm:"foreignkey:CidadeID"`

}