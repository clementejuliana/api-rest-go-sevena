package models

import "gorm.io/gorm"


type Administrador struct {
    gorm.Model
    Usuario   Usuario `gorm:"foreignKey:UsuarioID"`
    UsuarioID int     `json:"usuario_id,omitempty"`
    
}