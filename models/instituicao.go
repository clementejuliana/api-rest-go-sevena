package models

type Instituicao struct {
	InstituicaoID int `json:"instituicao_id,omitempty"`
	Status   string `json:"status,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Sigla    string `json:"sigla,omitempty"`
	CNPJ     string `json:"cnpj,omitempty"`
	Endereco string `json:"endereco,omitempty"`
	Telefone string `json:"telefone,omitempty"`
	Email    string `json:"email,omitempty"`
	CidadeID int    `json:"cidade_id,omitempty"`
}
