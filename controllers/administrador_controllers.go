package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirAdministrador(c *gin.Context) {
	var administrador []models.Administrador
	databasee.DB.Find(&administrador)
	c.JSON(200, administrador)
}

// criar esse novo administrador
func CriarNovoAdministrador(c *gin.Context) {
	var administrador models.Administrador
	if err := c.ShouldBindJSON(&administrador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&administrador)
	c.JSON(http.StatusOK, administrador)
}

func BuscarAdministradorPorID(c *gin.Context) {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.First(&administrador, id)

	if administrador.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Administrador Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, administrador)
}

func DeleteAdministrador(c *gin.Context) {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.Delete(&administrador, id)
	c.JSON(http.StatusOK, gin.H{"data": "Administrador deletado com sucesso"})
}

func EditarAdministrador(c *gin.Context)  {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.First(&administrador, id)

	if err := c.ShouldBindJSON(&administrador); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&administrador).UpdateColumns(administrador)
	c.JSON(http.StatusOK, administrador)
	
}
