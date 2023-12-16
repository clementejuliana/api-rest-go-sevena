package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirTipoUsuario(c *gin.Context) {
	var tipoUsuario []models.TipoUsuario
	databasee.DB.Find(&tipoUsuario)
	c.JSON(200, tipoUsuario)

}

//exibir uma mensagem quando está passando um valoe não valido
// func Saudacao(c *gin.Context) {
// 	nome := c.Params.ByName("nome")
// 	c.JSON(200, gin.H{
// 		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
// 	})
// }

// criar esse novo aluno
func CriarTipoUsuario(c *gin.Context) {
	var tipoUsuario models.TipoUsuario
	if err := c.ShouldBindJSON(&tipoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := tipoUsuario.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&tipoUsuario)
	c.JSON(http.StatusOK, tipoUsuario)
}

func BuscarTipoUsuarioPorID(c *gin.Context) {
	var tipoUsuario models.TipoUsuario
	id := c.Params.ByName("id")
	databasee.DB.First(&tipoUsuario, id)

	if tipoUsuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Tipo de Usuario não foi encontrando"})
		return
	}
	c.JSON(http.StatusOK, tipoUsuario)
}

func DeleteTipoUsuario(c *gin.Context) {
	var tipoUsuario models.TipoUsuario
	id := c.Params.ByName("id")
	databasee.DB.Delete(&tipoUsuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Tipo de Usuario deletado com sucesso"})
}

func EditarTipoUsuario(c *gin.Context) {
	var tipoUsuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&tipoUsuario, id)

	if err := c.ShouldBindJSON(&tipoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&tipoUsuario).UpdateColumns(tipoUsuario)
	c.JSON(http.StatusOK, tipoUsuario)

}

func BuscarTipoUsuarioPorCPF(c *gin.Context) {
	var tipoUsuario models.TipoUsuario
	cpf := c.Param("cpf")
	databasee.DB.Where(&models.Usuario{CPF: cpf}).First(&tipoUsuario)

	if tipoUsuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Tipo de Usuario não encontrando"})
		return
	}
	c.JSON(http.StatusOK, tipoUsuario)

}
