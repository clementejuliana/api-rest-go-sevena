package models

import "gorm.io/gorm"


type Administrador struct {
    gorm.Model
    Usuario   Usuario 
    UsuarioID int     `json:"usuario_id,omitempty"`
    
}