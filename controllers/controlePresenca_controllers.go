package controllers

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/services"
	"github.com/gin-gonic/gin"
	
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

func FiltrarControlePresenca(c *gin.Context) {
    var filtros models.ControlePresenca
    if err := c.BindJSON(&filtros); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de filtro inválidos"})
        return
    }

    var controlePresenca []models.ControlePresenca
    if err := databasee.DB.Where(&filtros).Find(&controlePresenca).Error; err != nil {
        // Trate o erro, por exemplo, retornando 500 Internal Server Error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar registros de presença"})
        return
    }

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


func GerarCertificado(c *gin.Context) {
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