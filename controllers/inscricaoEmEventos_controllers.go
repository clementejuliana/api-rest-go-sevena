package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos []models.InscricaoEmEvento
	databasee.DB.Find(&inscricaoEmEventos)
	c.JSON(200, inscricaoEmEventos)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoInscricaoEmEventos(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar essa nova inscricao em eventos
func CriarNovaInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	if err := c.ShouldBindJSON(&inscricaoEmEventos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := inscricaoEmEventos.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&inscricaoEmEventos)
	c.JSON(http.StatusOK, inscricaoEmEventos)
}

func BuscarInscricaoEmEventosPorID(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmEventos, id)

	if inscricaoEmEventos.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Inscrição em eventos Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, inscricaoEmEventos)
}

func DeleteInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.Delete(&inscricaoEmEventos, id)
	c.JSON(http.StatusOK, gin.H{"data": "Inscrição em eventos deletada com sucesso"})
}

func EditarInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmEventos, id)

	if err := c.ShouldBindJSON(&inscricaoEmEventos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&inscricaoEmEventos).UpdateColumns(inscricaoEmEventos)
	c.JSON(http.StatusOK, inscricaoEmEventos)

}
func GetInscritosNoEvento(c *gin.Context) {
    // Obtém o ID do evento da URL (por exemplo, /eventos/1/inscritos)
    eventoID := c.Param("eventoID")

    // Busca o evento no banco de dados
    evento := models.Evento{}
    if err := databasee.DB.Where("id = ?", eventoID).First(&evento).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Busca os inscritos no evento
    inscritos, err := evento.GetInscritos(databasee.DB)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Retorna os inscritos para o cliente
    c.JSON(200, inscritos)
}