package models

import "gorm.io/gorm"

type Notificacao struct {
	gorm.Model
	UsuarioID     int `gorm:"foreignKey:UsuarioID"`
	Conteudo string `json:"conteudo"`
}

