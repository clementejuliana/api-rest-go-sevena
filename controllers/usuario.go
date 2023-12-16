package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirUsuario(c *gin.Context) {
	var usuario []models.Usuario
	databasee.DB.Find(&usuario)
	c.JSON(200, usuario)

}

//exibir uma mensagem quando está passando um valoe não valido
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovoUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := usuario.Preparar(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	

	
	databasee.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func BuscarUsuarioPorID(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if usuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func DeleteUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario deletado com sucesso"})
}

func EditarUsuario(c *gin.Context)  {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)
	
}

func BuscarUsuarioPorCPF(c *gin.Context)  {
	var usuario models.Usuario
	cpf := c.Param("cpf")
	databasee.DB.Where(&models.Usuario{CPF: cpf}).First(&usuario)
	
	if usuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, usuario)

}
