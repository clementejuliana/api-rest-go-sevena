package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirNotificacao(c *gin.Context) {
	var notificacao []models.Notificacao
	databasee.DB.Find(&notificacao)
	c.JSON(200, notificacao)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoNotificacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar essa nova notificacao
func CriarNovaNotificacao(c *gin.Context) {
	var notificacao models.Notificacao
	if err := c.ShouldBindJSON(&notificacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&notificacao)
	c.JSON(http.StatusOK, notificacao)
}

func BuscarNotificacaoPorID(c *gin.Context) {
	var notificacao models.Notificacao
	id := c.Params.ByName("id")
	databasee.DB.First(&notificacao, id)

	if notificacao.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Notificação Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, notificacao)
}

func DeleteNotificacao(c *gin.Context) {
	var notificacao models.Notificacao
	id := c.Params.ByName("id")
	databasee.DB.Delete(&notificacao, id)
	c.JSON(http.StatusOK, gin.H{"data": "Notificação deletada com sucesso"})
}

func EditarNotificacao(c *gin.Context)  {
	var notificacao models.Notificacao
	id := c.Params.ByName("id")
	databasee.DB.First(&notificacao, id)

	if err := c.ShouldBindJSON(&notificacao); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&notificacao).UpdateColumns(notificacao)
	c.JSON(http.StatusOK, notificacao)
	
}