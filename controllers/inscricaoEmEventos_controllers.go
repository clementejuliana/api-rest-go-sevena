package controllers

import (
	"fmt"
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos []models.InscricaoEmEvento
	databasee.DB.Find(&inscricaoEmEventos)
	c.JSON(200, inscricaoEmEventos)
}

// criar essa nova inscricao em eventos
func CriarNovaInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	if err := c.ShouldBindJSON(&inscricaoEmEventos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := inscricaoEmEventos.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&inscricaoEmEventos)
	c.JSON(http.StatusOK, inscricaoEmEventos)
}

func BuscarInscricaoEmEventosPorID(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmEventos, id)

	if inscricaoEmEventos.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Inscrição em eventos Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, inscricaoEmEventos)
}

func DeleteInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.Delete(&inscricaoEmEventos, id)
	c.JSON(http.StatusOK, gin.H{"data": "Inscrição em eventos deletada com sucesso"})
}

func EditarInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos models.InscricaoEmEvento
	id := c.Params.ByName("id")
	databasee.DB.First(&inscricaoEmEventos, id)

	if err := c.ShouldBindJSON(&inscricaoEmEventos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&inscricaoEmEventos).UpdateColumns(inscricaoEmEventos)
	c.JSON(http.StatusOK, inscricaoEmEventos)

}

type localDetails struct {
	Setor string `json:"setor"`
	Sala  string  `json:"sala"`
}

func GetInscritosNoEvento(c *gin.Context) {
	// Obtém o ID do evento da URL (por exemplo, /eventos/1/inscritos)
	eventoID := c.Param("id")

	// Busca o evento no banco de dados
	evento := models.Evento{}
	if err := databasee.DB.Where("id = ?", eventoID).First(&evento).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Busca os inscritos no evento
	inscritos, err := evento.GetInscritos(databasee.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Define o cabeçalho "Content-Disposition" para que o arquivo seja baixado com o nome "inscritos.csv"
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=inscritosnoEvento.csv")

	// Define o cabeçalho "Content-Type" para que o arquivo seja salvo como um arquivo CSV
	c.Writer.Header().Set("Content-Type", "text/csv")

	// Escreve os cabeçalhos do arquivo CSV
	fmt.Fprintf(c.Writer, "Nome,Nome do evento,Data,Horário\n")

	// Escreve os dados dos inscritos no arquivo CSV
	for _, inscrito := range inscritos {
		
		fmt.Fprintf(c.Writer, "%s,%s,%s,%s\n",
			inscrito.Nome, evento.Nome, evento.DataInicio.Format("01-01-2006"), evento.HoraInicio.Format("15:00"))
	}

	// Retorna os inscritos para o cliente
	//  c.JSON(200, inscritos)
}
