package models

import "gorm.io/gorm"

type TipoUsuario struct {
	gorm.Model
	Status string `json:"status,omitempty"`
	TipoUsuario string `json:"tipo_usuario,omitempty"`
	Usuarios []Usuario `gorm:"foreignkey:TipoUsuarioID"`
}