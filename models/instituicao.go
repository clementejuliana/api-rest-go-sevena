package models

import (
	"gorm.io/gorm"
)

type Instituicao struct {
	gorm.Model
	Status   string   `json:"status,omitempty"`
	Nome     string   `json:"nome,omitempty"`
	Sigla    string   `json:"sigla,omitempty"`
	CNPJ     string   `json:"cnpj,omitempty"`
	Endereco string   `json:"endereco"`
	Telefone string   `json:"telefone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Usuarios []Usuario `gorm:"many2many:instituicao_usuarios;"`
	Cidades  []Cidade `gorm:"foreignkey:InstituicaoID"`
}

