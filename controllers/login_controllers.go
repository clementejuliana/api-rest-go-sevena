package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirLogin(c *gin.Context) {
	var login []models.Login
	databasee.DB.Find(&login)
	c.JSON(200, login)
}

//exibir uma mensagem quando está passando um valoe não valido
func Saudacaos(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo login
func CriarNovoLogin(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&login)
	c.JSON(http.StatusOK, login)
}

// func BuscarLoginPorID(c *gin.Context) {
// 	var login models.Login
// 	id := c.Params.ByName("id")
// 	databasee.DB.First(&login, id)

// 	if login.ID == 00 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"Not found": "Login Não encontrado"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, login)
// }

func DeleteLogin(c *gin.Context) {
	var login models.Login
	id := c.Params.ByName("id")
	databasee.DB.Delete(&login, id)
	c.JSON(http.StatusOK, gin.H{"data": "Login deletado com sucesso"})
}

func EditarLogin(c *gin.Context)  {
	var login models.Login
	id := c.Params.ByName("id")
	databasee.DB.First(&login, id)

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&login).UpdateColumns(login)
	c.JSON(http.StatusOK, login)
}