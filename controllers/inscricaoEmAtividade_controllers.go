package controllers
import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade []models.InscricaoEmAtividade
	databasee.DB.Find(&inscricaoEmAtividade)
	c.JSON(200, inscricaoEmAtividade)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoInscricaoEmAtividade(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar essa nova inscricao em atividade
func CriarNovaInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	if err := c.ShouldBindJSON(&inscricaoEmAtividade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&inscricaoEmAtividade)
	c.JSON(http.StatusOK, inscricaoEmAtividade)
}

func BuscarInscricaoEmAtividadePorID(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmAtividade, id)

	if inscricaoEmAtividade.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Inscrição em atividade Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, inscricaoEmAtividade)
}

func DeleteInscricaoEmAtividade(c *gin.Context) {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.Delete(&inscricaoEmAtividade, id)
	c.JSON(http.StatusOK, gin.H{"data": "Inscrição em atividade deletada com sucesso"})
}

func EditarInscricaoEmAtividade(c *gin.Context)  {
	var inscricaoEmAtividade models.InscricaoEmAtividade
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmAtividade, id)

	if err := c.ShouldBindJSON(&inscricaoEmAtividade); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&inscricaoEmAtividade).UpdateColumns(inscricaoEmAtividade)
	c.JSON(http.StatusOK, inscricaoEmAtividade)
	
}