package models

import "gorm.io/gorm"

type Notificacao struct {
	gorm.Model
	UsuarioID     int `gorm:"foreignKey:UsuarioID"`
	NotificacaoID int `json:"notificacao_id"`
	Conteudo string `json:"conteudo"`
}

var Notificacaos []Notificacao
