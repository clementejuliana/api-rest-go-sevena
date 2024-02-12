package models

import (
	"errors"
	
	"regexp"
	

	"gorm.io/gorm"
)

type InscricaoEmEvento struct {
	gorm.Model
	Status    string `json:"status,omitempty"`
	Data      string `json:"data,omitempty"`
	Hora      string `json:"hora,omitempty"`
	EventoID  int    `json:"evento_id" gorm:"foreignKey:EventoID"`
	UsuarioID uint   `json:"usuario_id" gorm:"foreignKey:UsuarioID"`
}

func (inscricaoEvento *InscricaoEmEvento) Preparar() error {
	// Chama a função
	err := inscricaoEvento.ValidarInscricaoEvento()
	// Verifica se houve erros
	if err != nil {
		return err
	}

	// Retorna nil se não houver erros
	return nil
}

func (i *InscricaoEmEvento) ValidarInscricaoEvento() error {
	// Verifica se o status é válido
	if i.Status != "ativo" && i.Status != "inativo" {
		return errors.New("status é obrigatório e deve ser 'ativo' ou 'inativo'")
	}

	// Verifica se a data e hora são válidas
	if i.Data == "" || i.Hora == "" {
		return errors.New("data e hora são obrigatórias")
	}

	// Validação para aceitar apenas números na data
	if match, _ := regexp.MatchString("[0-9]{2}/[0-9]{2}/[0-9]{4}$", i.Data); !match {
		return errors.New("data deve conter apenas números")
	}

	// Validação para aceitar apenas números na hora
	if match, _ := regexp.MatchString("[0-9]{2}:[0-9]{2}$", i.Hora); !match {
		return errors.New("hora deve conter apenas números")
	}
	
	// Verifica se o ID do evento é válido
	if i.EventoID == 0 {
		return errors.New("evento é obrigatório")
	}

	// Verifica se o ID do usuário é válido
	if i.UsuarioID == 0 {
		return errors.New("usuario é obrigatório")
	}

	return nil
}

func (e *Evento) GetInscritos(db *gorm.DB) ([]Usuario, error) {
	var usuarios []Usuario

	// Busca todas as inscrições para o evento atual
	inscricoes := []InscricaoEmEvento{}
	if err := db.Where("evento_id = ?", e.ID).Find(&inscricoes).Error; err != nil {
		return nil, err
	}

	// Para cada inscrição, busca o usuário correspondente
	for _, inscricao := range inscricoes {
		usuario := Usuario{}
		if err := db.Where("id = ?", inscricao.UsuarioID).First(&usuario).Error; err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}


