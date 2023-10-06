package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExibirEstado(c *gin.Context) {
	var estados []models.Estado
	databasee.DB.Find(&estados)
	c.JSON(200, estados)
}

func GetEstados(c *gin.Context) {
	// Retorna todos os estados
	var estados []models.Estado
	databasee := c.MustGet("db").(*gorm.DB)
	databasee.Find(&estados)
	// Retorna os estados
	c.JSON(200, estados)
}

func GetEstadosPaginated(c *gin.Context) {
	// Pagina os estados
	var estados []models.Estado
	databasee := c.MustGet("db").(*gorm.DB)
	databasee.Limit(10).Offset(0).Find(&estados)
	// Retorna os estados
	c.JSON(200, estados)
}

func GetEstadoOrdem(c *gin.Context) {
	// Ordena os estados pelo nome
	var estados []models.Estado
	databasee := c.MustGet("db").(*gorm.DB)
	databasee.Order("nome").Find(&estados)

	// Retorna os estados
	c.JSON(200, estados)
}

func PesquisaEstadosNome(c *gin.Context) {
	// Pesquisa os estados pelo nome
	var estados []models.Estado
	databasee := c.MustGet("db").(*gorm.DB)
	databasee.Where("nome LIKE ?", "%A%").Find(&estados)

	// Retorna os estados
	c.JSON(200, estados)
}

func ExportEstados(c *gin.Context) {
	// Exporta os estados em um arquivo CSV
	var estados []models.Estado
	databasee := c.MustGet("db").(*gorm.DB)
	databasee.Find(&estados)

	// Converte os estados em um array de bytes
	bytes, err := json.Marshal(estados)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Seta o cabeçalho do arquivo
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=estados.csv")

	// Retorna os dados do arquivo
	c.Data(200,"text/csv", bytes)
}

func LoadStates(c *gin.Context) {
	// Obtém a lista de estados do IBGE
	resp, err := http.Get("https://servicodados.ibge.gov.br/api/v1/localidades/estados")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Lê os dados da resposta
	var estados []models.Estado
	if err := json.NewDecoder(resp.Body).Decode(&estados); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Salva os estados no banco de dados
	databasee := c.MustGet("db").(*gorm.DB)
	for _, estado := range estados {
		databasee.Create(&estado)
	}

	// Retorna um status de sucesso
	c.JSON(200, gin.H{"status": "success"})
}
