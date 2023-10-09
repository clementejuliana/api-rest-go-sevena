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

func EditarEstado(c *gin.Context)  {
	var estado models.Estado
	id := c.Params.ByName("id")
	databasee.DB.First(&estado, id)

	if err := c.ShouldBindJSON(&estado); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&estado).UpdateColumns(estado)
	c.JSON(http.StatusOK, estado)
	
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
