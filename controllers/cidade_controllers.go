package controllers

import (
	"encoding/json"
	"io/ioutil"
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

func CriarCidade(c *gin.Context) {
	var cidade models.Cidade
	if err := c.ShouldBindJSON(&cidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := cidade.Preparar(); err != nil {
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

func EditarCidade(c *gin.Context) {
	var cidade models.Cidade
	id := c.Params.ByName("id")
	databasee.DB.First(&cidade, id)

	if err := c.ShouldBindJSON(&cidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&cidade).UpdateColumns(cidade)
	c.JSON(http.StatusOK, cidade)

}

func CarregarCidades(c *gin.Context) {
	// Ler o arquivo JSON
	dados, err := ioutil.ReadFile("data/estados-cidades2.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Decodificar o arquivo JSON
	cidades := []models.Cidade{}
	if err := json.Unmarshal(dados, &cidades); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Salvar as cidades no banco de dados
	for _, cidade := range cidades {
		databasee.DB.Create(&cidade)
	}

	// Retornar as cidades
	c.JSON(http.StatusOK, cidades)
}
