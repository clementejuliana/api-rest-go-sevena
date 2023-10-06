package models

import (
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Status string `json:"status,omitempty"`
	Nome string `json:"nome,omitempty"`
	CPF string `json:"cpf,omitempty"`
	RG string `json:"rg,omitempty"`
	Genero string `json:"genero"`
	DataNascimento time.Time `json:"data_nascimento"`
	Email string `json:"email,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Escolaridade string `json:"escolaridade"`
	Profissao string `json:"profissao"`
	FotoPerfil string `json:"foto_perfil,omitempty"`
	TipoUsuarioID int `json:"tipo_usuario_id,omitempty"`
	InstituicaoID int `json:"instituicao_id,omitempty"`
	CidadeID int `json:"cidadeid,omitempty"`
	Instituicoes []Instituicao `gorm:"many2many:instituicao_usuarios;"`
	

}
// criar uma lista de usuario que é referente ao tipo

//var Usuarios []Usuario