package controllers

import (
	"bytes"
	//"encoding/json"
	"fmt"
	//"log"
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade []models.InscricaoEmAtividade
	databasee.DB.Find(&inscricaoEmAtividade)
	c.JSON(200, inscricaoEmAtividade)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoInscricaoEmAtividade(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar essa nova inscricao em atividade
func CriarNovaInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	if err := c.ShouldBindJSON(&inscricaoEmAtividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := inscricaoEmAtividade.Preparar(databasee.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	databasee.DB.Create(&inscricaoEmAtividade)
	c.JSON(http.StatusOK, inscricaoEmAtividade)
}

func BuscarInscricaoEmAtividadePorID(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmAtividade, id)

	if inscricaoEmAtividade.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Inscrição em atividade Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, inscricaoEmAtividade)
}

func DeleteInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.Delete(&inscricaoEmAtividade, id)
	c.JSON(http.StatusOK, gin.H{"data": "Inscrição em atividade deletada com sucesso"})
}

func EditarInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmAtividade, id)

	if err := c.ShouldBindJSON(&inscricaoEmAtividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&inscricaoEmAtividade).UpdateColumns(inscricaoEmAtividade)
	c.JSON(http.StatusOK, inscricaoEmAtividade)

}

type UsuarioResumido struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func RelatorioInscritosEmAtividade(c *gin.Context) {
	// Obter o ID da atividade dos parâmetros da rota
	atividadeID := c.Params.ByName("id")

	// Buscar os detalhes da atividade no banco de dados
	var atividade models.Atividade
	databasee.DB.First(&atividade, atividadeID)

	// Buscar as inscrições na atividade no banco de dados
	var inscricoes []models.InscricaoEmAtividade
	databasee.DB.Where("id = ?", atividadeID).Find(&inscricoes)

	// Criar um slice para armazenar os detalhes dos usuários
	var usuariosResumidos []UsuarioResumido

	var local []models.Local
	databasee.DB.First(&local, atividadeID)
	// Para cada inscrição, buscar os detalhes do usuário e adicionar ao slice
	for _, inscricao := range inscricoes {
		var usuario models.Usuario
		databasee.DB.First(&usuario, inscricao.UsuarioID)

		// Criar um novo usuário resumido
		usuarioResumido := UsuarioResumido{
			Nome:  usuario.Nome,
			Email: usuario.Email,
		}

		// Adicionar o usuário resumido ao slice
		usuariosResumidos = append(usuariosResumidos, usuarioResumido)
	}

	// Cria um buffer para armazenar o relatório
	var buffer bytes.Buffer

	quantidadeInscritosFormatada := fmt.Sprintf("%d inscritos", atividade.QuantidadeInscritos)
	// Escreve o cabeçalho do relatório
	fmt.Fprintf(&buffer, "Relatório de inscrições em atividade\n\n")
	fmt.Fprintf(&buffer, "Atividade: %s\n\n", atividade.Titulo)
	fmt.Fprintf(&buffer, "Ministrante: %s\n\n", atividade.Ministrante)
	fmt.Fprintf(&buffer, "Data: %s\n\n", atividade.Data.Format("01-01-2006"))

	fmt.Fprintf(&buffer, "Horário: %s\n\n", atividade.HoraInicio.Format("15:00"))


	// Escrever a quantidade de inscritos no relatório
	fmt.Fprintf(&buffer, "Quantidade de inscritos: %s\n\n", quantidadeInscritosFormatada)
	// Escreve os detalhes dos usuários no relatório
	for _, usuario := range usuariosResumidos {
		fmt.Fprintf(&buffer, "Nome: %s, Email: %s\n", usuario.Nome, usuario.Email)
	}

	// Converte o buffer para uma string e retorna como resposta
	c.String(http.StatusOK, buffer.String())

	// // Serializar os detalhes dos usuários resumidos como um objeto JSON
	// usuariosResumidosJSON, err := json.Marshal(usuariosResumidos)
	// if err != nil {
	//     log.Println(err)
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//     return
	// }

	// Retornar os dados serializados
	//c.JSON(http.StatusOK, usuariosResumidosJSON)

	//c.JSON(http.StatusOK, gin.H{"relatorio": buffer.String(), "usuarios": usuariosResumidos})
	c.JSON(http.StatusOK, usuariosResumidos)
}
