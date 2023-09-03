package models

import "time"

type Usuario struct {
	UsuarioID uint64 `json:"usuario_id,omitempty"`
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
	CidadeID int `json:"cidade_id,omitempty"`

}