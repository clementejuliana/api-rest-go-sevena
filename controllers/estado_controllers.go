package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirEstados(c *gin.Context) {
	var estados []models.Estado
	databasee.DB.Find(&estados)
	c.JSON(200, estados)
}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoEstado(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovoEstado(c *gin.Context) {
	var estado models.Estado
	if err := c.ShouldBindJSON(&estado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := estado.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&estado)
	c.JSON(http.StatusOK, estado)
}

func BuscarEstadoPorID(c *gin.Context) {
	var estado models.Estado
	id := c.Params.ByName("id")
	databasee.DB.First(&estado, id)

	if estado.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Estado Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, estado)
}

func DeleteEstado(c *gin.Context) {
	var estado models.Estado
	id := c.Params.ByName("id")
	databasee.DB.Delete(&estado, id)
	c.JSON(http.StatusOK, gin.H{"data": "Estado deletado com sucesso"})
}

func EditarEstado(c *gin.Context) {
	var estado models.Estado
	id := c.Params.ByName("id")
	databasee.DB.First(&estado, id)

	if err := c.ShouldBindJSON(&estado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&estado).UpdateColumns(estado)
	c.JSON(http.StatusOK, estado)

}

func ExibirEstado(c *gin.Context) {
	var estados []models.Estado
	databasee.DB.Find(&estados)
	c.JSON(200, estados)
}

func CarregarEstados(c *gin.Context) {
	// Obter o link da IBGE
	link := "https://api.ibge.gov.br/v1/localidades/estados"

	// Fazer uma requisição GET ao link da IBGE
	response, err := http.Get(link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Decodificar a resposta da requisição
	var dados []models.Estado
	if err := json.NewDecoder(response.Body).Decode(&dados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Salvar os estados no banco de dados
	for _, estado := range dados {
		databasee.DB.Create(&estado)
	}

	// Retornar um status OK
	c.JSON(http.StatusOK, gin.H{"data": "Estados carregados com sucesso"})
}

