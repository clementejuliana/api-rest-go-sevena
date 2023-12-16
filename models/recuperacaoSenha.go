package models

import (
	"time"

	"gorm.io/gorm"
)

type RecuperacaoSenha struct {
	gorm.Model
	UsuarioID     int       `gorm:"foreignKey:UsuarioID"`
	Token         string    `json:"token"`
	Email     string `json:"email"`
	DataExpiracao time.Time `json:"data_expiracao"`
}

