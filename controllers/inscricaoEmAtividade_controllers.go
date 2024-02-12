package controllers

import (

	

	"fmt"
	"strconv"

	//"log"
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/services"
	"github.com/gin-gonic/gin"
)

func ExibirInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade []models.InscricaoEmAtividade
	databasee.DB.Find(&inscricaoEmAtividade)
	c.JSON(200, inscricaoEmAtividade)
}


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

	// Cria um novo controle de presença para a inscrição
	controlePresenca := models.ControlePresenca{
		Status: "Pendente",
		InscricaoAtividade: inscricaoEmAtividade.ID,
	}
	databasee.DB.Create(&controlePresenca)

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

	var local []models.Local
	if err := databasee.DB.First(&local, atividade.LocalID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//databasee.DB.First(&local, atividadeID)

	// Criar um slice para armazenar os detalhes dos usuários
	var usuariosResumidos []UsuarioResumido

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

	// Define o cabeçalho "Content-Disposition" para que o arquivo seja baixado com o nome "relatorio.csv"
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=relatoriodosinscritosporatividades.csv")

	// Define o cabeçalho "Content-Type" para que o arquivo seja salvo como um arquivo CSV
	c.Writer.Header().Set("Content-Type", "text/csv")
	quantidadeInscritosFormatada := fmt.Sprintf("%d inscritos", atividade.QuantidadeInscritos)

	// Escreve os cabeçalhos do arquivo CSV
	fmt.Fprintf(c.Writer, "Atividade:,Ministrante:,Data:,Hora Inicio:,Quantidade dos Inscritos:,Setor:,Sala:\n")
	// Escreve os detalhes da atividade no arquivo CSV
	fmt.Fprintf(c.Writer, "%s,%s,%s,%s,%s,%s,%s\n",
    atividade.Titulo, atividade.Ministrante, atividade.Data,
    atividade.HoraInicio,quantidadeInscritosFormatada,atividade.Local.Setor, atividade.Local.Sala)

	// Escreve os cabeçalhos dos usuários no arquivo CSV
	fmt.Fprintf(c.Writer, "\nDetalhes dos Inscritos\n")
	fmt.Fprintf(c.Writer, "Nome,Email\n")
	// Escreve os detalhes dos usuários no arquivo CSV
	for _, usuario := range usuariosResumidos {
		fmt.Fprintf(c.Writer, "%s,%s\n", usuario.Nome, usuario.Email)
	}
}

func ListarInscritosController(c *gin.Context) {
	atividadeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de atividade inválido"})
		return
	}

	inscritos, err := services.ListarInscritos(databasee.DB, atividadeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inscritos)
}


func ListarPresencas(c *gin.Context) {
    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Where("status = ?", "Presente").Find(&controlePresenca).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, controlePresenca)
}


func ListarPresencaPendente(c *gin.Context) {
    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Where("status = ?", "Pendente").Find(&controlePresenca).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, controlePresenca)
}

