package controllers

import (
	
	"fmt"
	"net/http"

	"time"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func ExibirInscricaoEmEventos(c *gin.Context) {
	var inscricaoEmEventos []models.InscricaoEmEvento
	databasee.DB.Find(&inscricaoEmEventos)
	c.JSON(200, inscricaoEmEventos)
}

// Função para converter a data e a hora para o formato desejado
func ConverterDataHora(data string, hora string) (time.Time, error) {
	// Concatena a data e a hora
	dataHora := fmt.Sprintf("%sT%s:00Z07:00", data, hora)

	// Converte para o formato time.Time
	t, err := time.Parse("2023-10-10T21:00:00-03:00", dataHora)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
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
			inscrito.Nome, evento.Nome, evento.DataInicio, evento.HoraInicio)
	}

	// Retorna os inscritos para o cliente
	//  c.JSON(200, inscritos)
}


func GerarCertificado(c *gin.Context) {
    // Obtém o ID do evento da URL
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

    // Para cada inscrito, gera um certificado em PDF
    for _, inscrito := range inscritos {
        atividade := models.Atividade{} // Substitua isso pela lógica correta para obter a atividade do inscrito
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
	
		// Configuração de fonte e estilo
		pdf.SetFont("Arial", "B", 20)
		
		// Título do certificado
		pdf.Cell(0, 10, "Certificado de Participação do Evento")
		pdf.Ln(20) // Espaçamento entre linhas
	
		// Configuração de fonte e estilo para informações detalhadas
		pdf.SetFont("Arial", "", 16)
	
		// Informações do certificado
		pdf.Cell(0, 10, "Certificamos que : " + inscrito.Nome)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Participou do Evento: " + evento.Nome)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Participando da Atividade: " + atividade.Titulo)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Data do Evento: " + evento.DataInicio)
		pdf.Ln(10)
		pdf.Cell(0, 10, "Carga Horária: " + fmt.Sprintf("%d", atividade.CargaHoraria) + " horas")
	
	

        // Salva o PDF
        err := pdf.OutputFileAndClose("certificado_" + inscrito.Nome + ".pdf")
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(200, gin.H{"message": "Certificados gerados com sucesso"})
}


