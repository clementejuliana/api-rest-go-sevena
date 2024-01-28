package services

import (
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"gorm.io/gorm"
)

func ListarInscritos(db *gorm.DB, atividadeID int) ([]models.InscricaoEmAtividade, error) {
	var inscricoes []models.InscricaoEmAtividade

	// Buscar todas as inscrições para a atividade
	if err := db.Find(&inscricoes, "atividade_id = ?", atividadeID).Error; err != nil {
		return nil, err
	}

	return inscricoes, nil
}