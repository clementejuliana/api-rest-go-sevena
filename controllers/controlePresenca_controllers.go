package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirControlePresenca(c *gin.Context) {
    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Find(&controlePresenca).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, controlePresenca)
}


func CriarNovoControlePresenca(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	if err := c.ShouldBindJSON(&controlePresenca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := controlePresenca.Preparar(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&controlePresenca)
	c.JSON(http.StatusOK, controlePresenca)
}


func DeleteControlePresenca(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.Delete(&controlePresenca, id)
	c.JSON(http.StatusOK, gin.H{"data": "Controle de Presença deletado com sucesso"})
}

func EditarControlePresenca(c *gin.Context)  {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.First(&controlePresenca, id)

	if err := c.ShouldBindJSON(&controlePresenca); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&controlePresenca).UpdateColumns(controlePresenca)
	c.JSON(http.StatusOK, controlePresenca)
	
}

func FiltrarControlePresenca(c *gin.Context) {
    var filtros models.ControlePresenca
    if err := c.BindJSON(&filtros); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de filtro inválidos"})
        return
    }

    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Where(&filtros).Find(&controlePresenca).Error; err != nil {
        // Trate o erro, por exemplo, retornando 500 Internal Server Error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar registros de presença"})
        return
    }

    c.JSON(http.StatusOK, controlePresenca)
}
