package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirLocal(c *gin.Context) {
	var local []models.Local
	databasee.DB.Find(&local)
	c.JSON(200, local)

}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoLocal(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovoLocal(c *gin.Context) {
	var local models.Local
	if err := c.ShouldBindJSON(&local); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := local.Preparar(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&local)
	c.JSON(http.StatusOK, local)
}

func BuscarLocalPorID(c *gin.Context) {
	var local models.Local
	id := c.Params.ByName("id")
	databasee.DB.First(&local, id)

	if local.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Local Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, local)
}

func DeleteLocal(c *gin.Context) {
	var local models.Local
	id := c.Params.ByName("id")
	databasee.DB.Delete(&local, id)
	c.JSON(http.StatusOK, gin.H{"data": "Local deletado com sucesso"})
}

func EditarLocal(c *gin.Context)  {
	var local models.Local
	id := c.Params.ByName("id")
	databasee.DB.First(&local, id)

	if err := c.ShouldBindJSON(&local); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&local).UpdateColumns(local)
	c.JSON(http.StatusOK, local)
	
}

// func BuscarLocalPorCPF(c *gin.Context)  {
// 	var local models.Local
// 	cpf := c.Param("cpf")
// 	databasee.DB.Where(&models.Local{CPF: cpf}).First(&local)
	
// 	if local.ID == 00 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"Not found": "Local Não encontrando"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, local)

// }