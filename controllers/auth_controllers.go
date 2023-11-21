package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirAuth(c *gin.Context) {
	var auth []models.Auth
	databasee.DB.Find(&auth)
	c.JSON(200, auth)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoAuth(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovaAuth(c *gin.Context) {
	var auth models.Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&auth)
	c.JSON(http.StatusOK, auth)
}

func BuscarAuthPorID(c *gin.Context) {
	var auth models.Auth
	id := c.Params.ByName("id")
	databasee.DB.First(&auth, id)

	if auth.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Auth Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, auth)
}

func DeleteAuth(c *gin.Context) {
	var auth models.Auth
	id := c.Params.ByName("id")
	databasee.DB.Delete(&auth, id)
	c.JSON(http.StatusOK, gin.H{"data": "Auth deletada com sucesso"})
}

func EditarAuth(c *gin.Context)  {
	var auth models.Auth
	id := c.Params.ByName("id")
	databasee.DB.First(&auth, id)

	if err := c.ShouldBindJSON(&auth); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&auth).UpdateColumns(auth)
	c.JSON(http.StatusOK, auth)
	
}