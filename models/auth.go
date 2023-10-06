package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
    gorm.Model
	UsuarioID int `json:"usuario_id"`
    Token string `json:"token"`
    DataExpiracao time.Time `json:"data_expiracao"`
}