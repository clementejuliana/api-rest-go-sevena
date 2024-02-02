package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirRecuperacaoSenha(c *gin.Context) {
	var recuperacaoSenha []models.RecuperacaoSenha
	databasee.DB.Find(&recuperacaoSenha)
	c.JSON(200, recuperacaoSenha)
}

// criar esse novo recuperacaoSenha
func CriarNovoRecuperacaoSenha(c *gin.Context) {
	var recuperacaoSenha models.RecuperacaoSenha
	if err := c.ShouldBindJSON(&recuperacaoSenha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&recuperacaoSenha)
	c.JSON(http.StatusOK, recuperacaoSenha)
}

func BuscarRecuperacaoSenhaPorID(c *gin.Context) {
	var recuperacaoSenha models.RecuperacaoSenha
	id := c.Params.ByName("id")
	databasee.DB.First(&recuperacaoSenha, id)

	if recuperacaoSenha.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "RecuperacaoSenha Não encontrada"})
		return
	}
	c.JSON(http.StatusOK, recuperacaoSenha)
}

func DeleteRecuperacaoSenha(c *gin.Context) {
	var recuperacaoSenha models.RecuperacaoSenha
	id := c.Params.ByName("id")
	databasee.DB.Delete(&recuperacaoSenha, id)
	c.JSON(http.StatusOK, gin.H{"data": "RecuperacaoSenha deletada com sucesso"})
}

func EditarRecuperacaoSenha(c *gin.Context)  {
	var recuperacaoSenha models.RecuperacaoSenha
	id := c.Params.ByName("id")
	databasee.DB.First(&recuperacaoSenha, id)

	if err := c.ShouldBindJSON(&recuperacaoSenha); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&recuperacaoSenha).UpdateColumns(recuperacaoSenha)
	c.JSON(http.StatusOK, recuperacaoSenha)
	
}