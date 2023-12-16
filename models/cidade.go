package models

import (
	"errors"

	"gorm.io/gorm"
)

type Cidade struct {
	gorm.Model
	Status   string `json:"status,omitempty"`
	Nome string`json:"nome,omitempty"`
	InstituicaoID int
	EstadoID int `json:"estado_id,omitempty"`
	Usuarios []Usuario `gorm:"foreignkey:CidadeID"`

}
func (cidade *Cidade) Preparar() error {
	// Chama a função ValidarUsuario()
	err := cidade.ValidarCidade()
	// Verifica se houve erros
	if err != nil {
		return err
	}
	// Retorna nil se não houver erros
	return nil
}
func (cidade *Cidade) ValidarCidade() error {
    // Valida o status
    if cidade.Status != "ativo" && cidade.Status != "inativo" {
        return errors.New("status inválido")
    }
    
    // Valida o nome
    if len(cidade.Nome) < 3 {
        return errors.New("nome deve ter pelo menos 3 caracteres")
    }
    
    // Valida o InstituicaoID
    if cidade.InstituicaoID < 1 {
        return errors.New("ID da instituição inválido")
    }
    
    // Valida o EstadoID
    if cidade.EstadoID < 1 {
        return errors.New("ID do estado inválido")
    }
    
    return nil
}