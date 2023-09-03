package models

type Notificacao struct {
    NotificacaoID int `json:"notificacao_id"`
    UsuarioID int `json:"usuario_id"`
    Conteudo string `json:"conteudo"`
}

var Notificacaos []Notificacao