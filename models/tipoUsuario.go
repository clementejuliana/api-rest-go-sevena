package models

type TipoUsuario struct {
	TipoUsuarioID int `json:"tipo_usuario_id,omitempty"`
	Status string `json:"status,omitempty"`
	TipoUsuario string `json:"tipo_usuario,omitempty"`
}