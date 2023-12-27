package models

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)


type Administrador struct {
    gorm.Model
    Usuario   Usuario 
    UsuarioID int     `json:"usuario_id,omitempty"`
	Status             string    `json:"status,omitempty"`
    
}


func ValidateAdministrador(administrador Administrador) error {
	err := validator.New().Struct(administrador)
	if err != nil {
		return err
	}
	return nil
}