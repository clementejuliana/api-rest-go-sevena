package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/services"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func ExibirControlePresenca(c *gin.Context) {
    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Find(&controlePresenca).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.JSON(http.StatusOK, controlePresenca)
}


func CriarNovoControlePresenca(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	if err := c.ShouldBindJSON(&controlePresenca); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := controlePresenca.Preparar(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&controlePresenca)
	c.JSON(http.StatusOK, controlePresenca)
}


func DeleteControlePresenca(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.Delete(&controlePresenca, id)
	c.JSON(http.StatusOK, gin.H{"data": "Controle de Presença deletado com sucesso"})
}

func EditarControlePresenca(c *gin.Context)  {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.First(&controlePresenca, id)

	if err := c.ShouldBindJSON(&controlePresenca); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&controlePresenca).UpdateColumns(controlePresenca)
	c.JSON(http.StatusOK, controlePresenca)
	
}

func RegistrarPresencak(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	status := c.Query("status") // Adicione esta linha

	if status != "Presente" && status != "Ausente" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status inválido. Deve ser 'Presente' ou 'Ausente'."})
		return
	}

	databasee.DB.First(&controlePresenca, id)

	controlePresenca.Status = status
	if status == "Presente" {
		controlePresenca.HoraEntrada = time.Now()
	} else {
		controlePresenca.HoraSaida = time.Now()
	}

	databasee.DB.Save(&controlePresenca)

	c.JSON(http.StatusOK, gin.H{"status": "Presença registrada com sucesso"})
}


func GerarCertificado8(c *gin.Context) {
	var controlePresenca models.ControlePresenca
	id := c.Params.ByName("id")
	databasee.DB.First(&controlePresenca, id)

	if controlePresenca.Status != "Presente" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Não é possível gerar certificado, presença não registrada"})
		return
	}

	// Aqui você pode adicionar o código para gerar o certificado.
	// Isso dependerá de como você deseja que os certificados sejam formatados e entregues.

	c.JSON(http.StatusOK, gin.H{"status": "Certificado gerado com sucesso"})
}

func GerarCertificadooo(c *gin.Context) {
	inscricaoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de inscrição inválido"})
		return
	}

	var controle models.ControlePresenca
	databasee.DB.First(&controle, inscricaoID)

	if controle.Status != "Presente" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Não é possível gerar certificado, presença não registrada"})
		return
	}

	// Gerar o certificado
	//certificado, err := services.GerarCertificadoPDF(controle.InscricaoEmAtividade.Usuario.Nome, controle.InscricaoEmAtividade.Atividade)
	// certificado, err := services.GerarCertificadoPDF(controle.InscricaoEmAtividade.Usuario.Nome, controle.InscricaoEmAtividade.Atividade.Titulo)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	if controle.InscricaoEmAtividade == nil || controle.InscricaoEmAtividade.Usuario == nil || controle.InscricaoEmAtividade.Atividade == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "InscricaoEmAtividade, Usuario ou Atividade é nil"})
		return
	}
	
	nome := controle.InscricaoEmAtividade.Usuario.Nome
	titulo := controle.InscricaoEmAtividade.Atividade.Titulo
	
	if nome == "" || titulo == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nome ou Titulo está vazio"})
		return
	}
	
	certificado, err := services.GerarCertificadoPDF(nome, titulo)
	

	// Salvar o certificado em um local temporário
	caminhoTemp := "/tmp/certificado_" + strconv.Itoa(inscricaoID) + ".pdf"
	err = services.SalvarCertificado(certificado, caminhoTemp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Enviar o certificado por e-mail
	// err = services.EnviarCertificadoPorEmail(caminhoTemp, controle.Email)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Certificado gerado com sucesso"})
}


func RegistrarPresenca(c *gin.Context) {
	inscricaoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de inscrição inválido"})
		return
	}

	compareceu := c.PostForm("compareceu") == "true"

	var controle models.ControlePresenca

	// Buscar o controle de presença associado à inscrição
	if err := databasee.DB.First(&controle, "inscricao_atividade = ?", inscricaoID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Atualizar o status e as horas de entrada e saída conforme a presença
	if compareceu {
		controle.Status = "Presente"
		controle.HoraEntrada = time.Now()
		controle.HoraSaida = time.Now()
	} else {
		controle.Status = "Ausente"
	}

	// Salvar as alterações no banco de dados
	if err := databasee.DB.Save(&controle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presença atualizada com sucesso"})
}



func AtualizarPresencaParaTodos(c *gin.Context) {
	atividadeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de atividade inválido"})
		return
	}

	// Obter a lista de inscritos
	inscritos, err := services.ListarInscritos(databasee.DB, atividadeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Percorrer a lista de inscritos
	for _, inscricao := range inscritos {
		// Atualizar o contexto com os parâmetros da inscrição
		c.Set("id", inscricao.ID)
		c.Request.PostForm = make(url.Values)
		c.Request.PostForm.Add("compareceu", "true") // ou "false" se o usuário não compareceu

		// Chamar a função RegistrarPresenca
		RegistrarPresenca(c)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presença atualizada para todos os inscritos"})
}


func GerarCertificadoPresenca(c *gin.Context) {
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
		pdf.Cell(0, 10, "Data do Evento: " + evento.DataInicio.Format("02 de janeiro de 2006"))
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
