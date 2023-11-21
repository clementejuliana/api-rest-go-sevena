package controllers


import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirAtividade(c *gin.Context) {
	var atividade []models.Atividade
	databasee.DB.Find(&atividade)
	c.JSON(200, atividade)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoAtividade(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar essa nova atividade
func CriarNovaAtividade(c *gin.Context) {
	var atividade models.Atividade
	if err := c.ShouldBindJSON(&atividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&atividade)
	c.JSON(http.StatusOK, atividade)
}

func BuscarAtividadePorID(c *gin.Context) {
	var atividade models.Atividade
	id := c.Params.ByName("id")
	databasee.DB.First(&atividade, id)

	if atividade.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Atividade Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, atividade)
}

func DeleteAtividade(c *gin.Context) {
	var atividade models.Atividade
	id := c.Params.ByName("id")
	databasee.DB.Delete(&atividade, id)
	c.JSON(http.StatusOK, gin.H{"data": "Atividade deletada com sucesso"})
}

func EditarAtividade(c *gin.Context)  {
	var atividade models.Atividade
	id := c.Params.ByName("id")
	databasee.DB.First(&atividade, id)

	if err := c.ShouldBindJSON(&atividade); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&atividade).UpdateColumns(atividade)
	c.JSON(http.StatusOK, atividade)
	
}
