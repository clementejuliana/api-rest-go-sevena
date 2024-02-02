package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirInstuicao(c *gin.Context) {
	var instituicao []models.Instituicao
	databasee.DB.Find(&instituicao)
	c.JSON(200, instituicao)

}

// criar esse novo aluno
func CriarNovaInstituicao(c *gin.Context) {
	var instituicao models.Instituicao
	if err := c.ShouldBindJSON(&instituicao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := instituicao.Preparar(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&instituicao)
	c.JSON(http.StatusOK, instituicao)
}

func BuscarInstuicaoPorID(c *gin.Context) {
	var instituicao models.Instituicao
	id := c.Params.ByName("id")
	databasee.DB.First(&instituicao, id)

	if instituicao.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Instituicao Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, instituicao)
}

func DeleteInstituicao(c *gin.Context) {
	var instituicao models.Instituicao
	id := c.Params.ByName("id")
	databasee.DB.Delete(&instituicao, id)
	c.JSON(http.StatusOK, gin.H{"data": "Instituição deletada com sucesso"})
}

func EditarInstituicao(c *gin.Context)  {
	var instituicao models.Instituicao
	id := c.Params.ByName("id")
	databasee.DB.First(&instituicao, id)

	if err := c.ShouldBindJSON(&instituicao); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&instituicao).UpdateColumns(instituicao)
	c.JSON(http.StatusOK, instituicao)
	
}

func BuscarPorInstituicao(c *gin.Context)  {
	var instituicao models.Instituicao
	nome := c.Param("nome")
	databasee.DB.Where(&models.Instituicao{Nome: nome}).First(&instituicao)
	
	if instituicao.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Instituição de Ensino não encontranda"})
		return
	}
	c.JSON(http.StatusOK, instituicao)
}