package models

import (
	"errors"
	"regexp"

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

func (instituicao *Instituicao) Preparar() error {
	// Chama a função ValidarU
	err := instituicao.ValidarInstituicao()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}


func (instituicao *Instituicao) ValidarInstituicao() error {
    if instituicao.Status == "" {
        return errors.New("campo 'Status' é obrigatório")
    }

    if instituicao.Nome == "" {
        return errors.New("campo 'Nome' é obrigatório")
    }

    if instituicao.CNPJ == "" {
        return errors.New("campo 'CNPJ' é obrigatório")
    }

    if !regexp.MustCompile(`^[0-9]{18}$`).MatchString(instituicao.CNPJ) {
        return errors.New("campo 'CNPJ' deve ter 14 dígitos")
    }

    if instituicao.Endereco == "" {
        return errors.New("campo 'Endereco' é obrigatório")
    }

    return nil
}