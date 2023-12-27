package controllers


import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirTipoAtividade(c *gin.Context) {
	var tipoAtividade []models.TipoAtividade
	databasee.DB.Find(&tipoAtividade)
	c.JSON(200, tipoAtividade)

}
// criar esse novo aluno
func CriarNovoTipoAtividade(c *gin.Context) {
	var tipoAtividade models.TipoAtividade
	if err := c.ShouldBindJSON(&tipoAtividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := tipoAtividade.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&tipoAtividade)
	c.JSON(http.StatusOK, tipoAtividade)
}

func BuscarTipoAtividadePorID(c *gin.Context) {
	var tipoAtividade models.TipoAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&tipoAtividade, id)

	if tipoAtividade.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "TipoAtividade Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, tipoAtividade)
}

func DeleteTipoAtividade(c *gin.Context) {
	var tipoAtividade models.TipoAtividade
	id := c.Params.ByName("id")
	databasee.DB.Delete(&tipoAtividade, id)
	c.JSON(http.StatusOK, gin.H{"data": "TipoAtividade deletado com sucesso"})
}

func EditarTipoAtividade(c *gin.Context)  {
	var tipoAtividade models.TipoAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&tipoAtividade, id)

	if err := c.ShouldBindJSON(&tipoAtividade); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&tipoAtividade).UpdateColumns(tipoAtividade)
	c.JSON(http.StatusOK, tipoAtividade)
	
}









