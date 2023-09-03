package models

import "time"

type Auth struct {
	UsuarioID int `json:"usuario_id"`
    Token string `json:"token"`
    DataExpiracao time.Time `json:"data_expiracao"`
}