package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirControlePresenca(c *gin.Context) {
	var controlePresenca []models.ControlePresenca
	databasee.DB.Find(&controlePresenca)
	c.JSON(200, controlePresenca)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoControlePresenca(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovoControlePresenca(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	if err := c.ShouldBindJSON(&controlePresenca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&controlePresenca)
	c.JSON(http.StatusOK, controlePresenca)
}

func BuscarControlePresencaPorID(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.First(&controlePresenca, id)

	if controlePresenca.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Controle de Presença Não encontrando"})
		return
	}
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