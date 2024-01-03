package models

import (
	"bytes"

	"errors"
	"fmt"

	"time"

	"gorm.io/gorm"
)

type InscricaoEmAtividade struct {
	gorm.Model
	AtividadeID        int               `json:"atividade_id,omitempty"`
	EventoID           int               `json:"evento_id,omitempty"`
	Status             string            `json:"status,omitempty"`
	Data               time.Time         `json:"data,omitempty"`
	Hora               time.Time         `json:"hora,omitempty"`
	ControlePresencaID uint              `json:"controle_presenca_id,omitempty" gorm:"foreignKey:ControlePresenca"`
	ControlePresenca   *ControlePresenca `json:"controle_presenca,omitempty"`
	UsuarioID          int               `json:"usuario_id,omitempty"`
	Usuario            *Usuario          `json:"usuario,omitempty"`
	Atividade          *Atividade        `json:"atividade,omitempty"`
}

func (inscricaoA *InscricaoEmAtividade) Preparar(db *gorm.DB) error {
	// Chama a função ValidarInscricaoAtividade
	err := inscricaoA.ValidarInscricaoAtividade()
	if err != nil {
		return err
	}

	//Verifica conflito de horário passando a instância do banco de dados
	err = inscricaoA.VerificarConflitoHorario(db)
	if err != nil {
		return err
	}

	// Retorna nil se não houver erros
	return nil
}

func (i *InscricaoEmAtividade) ValidarInscricaoAtividade() error {
	// Verifica se o ID da atividade é válido
	if i.AtividadeID == 0 {
		return errors.New("atividade é obrigatório")
	}

	// Verifica se o ID do evento é válido
	if i.EventoID == 0 {
		return errors.New("evento é obrigatório")
	}

	// Verifica se o status é válido
	if i.Status != "pendente" && i.Status != "confirmada" && i.Status != "cancelada" {
		return errors.New("status deve ser 'pendente', 'confirmada' ou 'cancelada'")
	}

	// Verifica se a data é válida
	if i.Data.IsZero() {
		return errors.New("data é obrigatória")
	}

	// Verifica se a hora é válida
	if i.Hora.IsZero() {
		return errors.New("hora é obrigatória")
	}

	// Verifica se o controle de presença é válido
	if i.ControlePresencaID == 0 {
		return errors.New("controle presenca é obrigatório")
	}

	// Se chegamos até aqui, todos os campos são válidos
	return nil
}

func (i *InscricaoEmAtividade) VerificarConflitoHorario(db *gorm.DB) error {
	// Consulta para verificar conflito de horário
	var count int64
	result := db.Model(&InscricaoEmAtividade{}).
		Where("Evento_ID = ? AND Data = ? AND ((Hora >= ? AND Hora < ?) OR (Hora <= ? AND ? < Hora))",
			i.EventoID, i.Data, i.Hora, i.Hora.Add(time.Hour), i.Hora, i.Hora.Add(time.Hour)).
		Count(&count)

	if result.Error != nil {
		return result.Error
	}

	if count > 0 {
		return errors.New("conflito de horário com outra inscrição")
	}

	return nil
}

func (a *Atividade) GetInscritos(db *gorm.DB) ([]Usuario, error) {
	var usuarios []Usuario

	// Busca todas as inscrições para a atividade atual
	inscricoes := []InscricaoEmAtividade{}
	if err := db.Where("atividade_id = ?", a.ID).Find(&inscricoes).Error; err != nil {
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

func GerarRelatorioInscritosAtividades(db *gorm.DB) ([]byte, error) {
	// Consulta as inscrições em atividades
	inscricoes := []InscricaoEmAtividade{}
	if err := db.Find(&inscricoes).Error; err != nil {
		return nil, err
	}

	// Cria um buffer para o relatório
	buffer := bytes.NewBuffer([]byte{})

	// Escreve o cabeçalho do relatório
	buffer.WriteString("Atividade | Nome | Email\n")

	// Itera sobre as inscrições
	for _, inscricao := range inscricoes {
		// Obtém o nome e o email do usuário
		nome := inscricao.Usuario.Nome
		email := inscricao.Usuario.Email

		// Obtém o nome da atividade
		titulo := inscricao.Atividade.Titulo

		// Escreve a linha do relatório
		buffer.WriteString(fmt.Sprintf("%s | %s | %s\n", titulo, nome, email))
	}

	// Retorna o relatório
	return buffer.Bytes(), nil
}

func RelatorioInscritosPorAtividade0000(db *gorm.DB) ([]byte, error) {
	// Consulta para obter os inscritos por atividade
	var inscritos []struct {
		AtividadeID   int
		UsuarioID     int
		NomeUsuario   string
		Email         string
		NomeAtividade string
	}
	result := db.Model(&InscricaoEmAtividade{}).
		Select("inscricao_em_atividades.atividade_id, usuario.id as usuario_id, usuario.nome as nome_usuario, usuario.email, atividade.nome as nome_atividade").
		Joins("INNER JOIN usuario ON usuario.id = inscricao_em_atividades.usuario_id").
		Joins("INNER JOIN atividade ON atividade.id = inscricao_em_atividades.atividade_id").
		Scan(&inscritos)

	if result.Error != nil {
		return nil, result.Error
	}

	// Gera o relatório
	var buffer bytes.Buffer
	for _, inscrito := range inscritos {
		buffer.WriteString("Atividade: " + inscrito.NomeAtividade + "\n")
		buffer.WriteString("Usuário: " + inscrito.NomeUsuario + "\n")
		buffer.WriteString("E-mail: " + inscrito.Email + "\n")
		buffer.WriteString("\n")
	}

	return buffer.Bytes(), nil
}

func RelatorioInscritosPorAtividade(db *gorm.DB) ([]byte, error) {
	var inscritos []struct {
		AtividadeID int
		UsuarioID   int
		Nome        string
		Email       string
	}

	// Utilize o método Preload para pré-carregar a relação com o usuário (assumindo que a relação está definida no modelo).
	result := db.Model(&InscricaoEmAtividade{}).
		Preload("usuario").
		Select("inscricao_em_atividades.atividade_id, usuario.usuario_id, usuario.nome, usuario.email").
		Scan(&inscritos)

	if result.Error != nil {
		return nil, result.Error
	}

	var buffer bytes.Buffer
	for _, inscrito := range inscritos {
		buffer.WriteString("Atividade: " + fmt.Sprintf("%d", inscrito.AtividadeID) + "\n")
		buffer.WriteString("Usuário: " + inscrito.Nome + "\n")
		buffer.WriteString("E-mail: " + inscrito.Email + "\n")
		buffer.WriteString("-------------------------------\n")
	}

	return buffer.Bytes(), nil
}

func RelatorioInscritosPorAtividade1(db *gorm.DB) ([]byte, error) {
	var inscritos []struct {
		AtividadeID int
		UsuarioID   int
		Nome        string
		Email       string
	}
	result := db.Model(&InscricaoEmAtividade{}).
		Select("inscricao_em_atividades.atividade_id, Usuario.usuario_id, Usuario.nome, Usuario.email").
		Scan(&inscritos)

	if result.Error != nil {
		return nil, result.Error
	}

	var buffer bytes.Buffer
	for _, inscrito := range inscritos {
		buffer.WriteString("Atividade: " + fmt.Sprintf("%d", inscrito.AtividadeID) + "\n")
		buffer.WriteString("Usuário: " + inscrito.Nome + "\n")
		buffer.WriteString("E-mail: " + inscrito.Email + "\n")
	}

	return buffer.Bytes(), nil
}
