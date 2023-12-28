package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/services"
	"github.com/gin-gonic/gin"
	
)

func ExibirUsuario(c *gin.Context) {
	var usuario []models.Usuario
	databasee.DB.Find(&usuario)
	c.JSON(200, usuario)

}

// criar esse novo aluno
func CriarNovoUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := usuario.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var usuarioExistente models.Usuario
	email := databasee.DB.Where("email = ?", usuario.Email).Find(&usuarioExistente)
	if email.RowsAffected > 0 {
		c.JSON(400, gin.H{"MENSAGEM": "Esse email já está cadastrado no sistema!"})
		return
	}

	usuario.Senha = services.SHA256Encoder(usuario.Senha)
	databasee.DB.Create(&usuario)
	c.JSON(http.StatusOK, gin.H{
		"data": "Cadastro realizado com sucesso!"})
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

func EditarUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)

}

func BuscarUsuarioPorCPF(c *gin.Context) {
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
