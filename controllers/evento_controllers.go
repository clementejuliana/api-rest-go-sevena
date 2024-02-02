package controllers

import (
	"net/http"


	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirEventos(c *gin.Context) {
	var evento []models.Evento
	databasee.DB.Find(&evento)
	c.JSON(200, evento)

}


func FiltrarEventos(c *gin.Context) {
	var evento []models.Evento
	filtro := c.Query("filtro")

	if filtro != "" {
		databasee.DB.Where("nome LIKE ?", "%" + filtro + "%").Find(&evento)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nenhum filtro fornecido"})
		return
	}

	c.JSON(200, evento)
}

// criar esse novo aluno
func CriarNovoEvento(c *gin.Context) {
	var evento models.Evento
	if err := c.ShouldBindJSON(&evento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := evento.Preparar(databasee.DB); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&evento)
	c.JSON(http.StatusOK, evento)
}

func BuscarEventoPorID(c *gin.Context) {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.First(&evento, id)

	if evento.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Evento Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, evento)
}

func DeleteEvento(c *gin.Context) {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.Delete(&evento, id)
	c.JSON(http.StatusOK, gin.H{"data": "Evento deletado com sucesso"})
}

func EditarEvento(c *gin.Context)  {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.First(&evento, id)

	if err := c.ShouldBindJSON(&evento); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&evento).UpdateColumns(evento)
	c.JSON(http.StatusOK, evento)
	
}
