package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirCidade(c *gin.Context) {
	var cidade []models.Cidade
	databasee.DB.Find(&cidade)
	c.JSON(200, cidade)

}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoCidade(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarCidade(c *gin.Context) {
	var cidade models.Cidade
	if err := c.ShouldBindJSON(&cidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&cidade)
	c.JSON(http.StatusOK, cidade)
}

func BuscarCidadePorID(c *gin.Context) {
	var cidade models.Cidade
	id := c.Params.ByName("id")
	databasee.DB.First(&cidade, id)

	if cidade.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Cidade não encontranda"})
		return
	}
	c.JSON(http.StatusOK, cidade)
}

func DeleteCidade(c *gin.Context) {
	var cidade models.Cidade
	id := c.Params.ByName("id")
	databasee.DB.Delete(&cidade, id)
	c.JSON(http.StatusOK, gin.H{"data": "Cidade deletada com sucesso"})
}

func EditarCidade(c *gin.Context)  {
	var cidade models.Cidade
	id := c.Params.ByName("id")
	databasee.DB.First(&cidade, id)

	if err := c.ShouldBindJSON(&cidade); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&cidade).UpdateColumns(cidade)
	c.JSON(http.StatusOK, cidade)
	
}

// func BuscarUsuarioPorCPF(c *gin.Context)  {
// 	var usuario models.Usuario
// 	cpf := c.Param("cpf")
// 	databasee.DB.Where(&models.Usuario{CPF: cpf}).First(&usuario)
	
// 	if usuario.ID == 00 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"Not found": "Usuario Não encontrando"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, usuario)

// }
