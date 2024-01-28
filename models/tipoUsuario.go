package models

import (
	"errors"

	"gorm.io/gorm"
)

type TipoUsuario struct {
	gorm.Model
	Status      string    `json:"status,omitempty"`
	TipoUsuario string    `json:"tipo_usuario,omitempty" gorm:"unique"` // Tipo de usuário deve ser único
	Usuarios    []Usuario `gorm:"foreignkey:TipoUsuarioID"`
}

func (tipoUsuario *TipoUsuario) Preparar() error {
	// Chama a função ValidarUsuario()
	err := tipoUsuario.ValidarTipoUsuario()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}

func (tipoUsuario *TipoUsuario) ValidarTipoUsuario() error {
	if tipoUsuario.Status != "ativo" && tipoUsuario.Status != "inativo" {
		return errors.New("status é obrigatório")
	}
	if tipoUsuario.TipoUsuario == "" {
		return errors.New("tipo de usuário é obrigatório")
	}
	return nil
}
