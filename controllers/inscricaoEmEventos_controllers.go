package controllers

import (
	"errors"
	"net/http"


	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func GetInscritosNoEvento(c *gin.Context) {
    // Obtém o ID do evento da URL (por exemplo, /eventos/1/inscritos)
    eventoID := c.Param("eventoID")

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

    // Retorna os inscritos para o cliente
    c.JSON(200, inscritos)
}

// GerarRelatorioInscritosEvento gera um relatório com a lista de usuários inscritos em um evento
func GerarRelatorioInscritosEvento3(c *gin.Context) {
	eventoID := c.Params.ByName("evento_id")
    // eventoIDStr := c.Param("evento_id")
	// eventoID, err := strconv.Atoi(eventoIDStr)
    // if err != nil {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
    //     return
    // }
    
    // Consultar inscrições para o evento específico
    var inscricoes []models.InscricaoEmEvento
    if err := databasee.DB.Where("evento_id = ? ", eventoID).Find(&inscricoes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar inscrições"})
        return
    }

    // Consultar detalhes dos usuários inscritos
    var usuarios []models.Usuario
    if err := databasee.DB.Model(&models.Usuario{}).Joins("JOIN inscricao_em_evento ON inscricao_em_eventos.usuario_id = usuarios.id").Where("inscricao_em_eventos.evento_id = ?", eventoID).Find(&usuarios).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar detalhes dos usuários"})
            return
        }
    }

	

    // Criar relatório combinando dados de inscrições e usuários
    relatorio := make([]gin.H, len(inscricoes))
    for i, inscricao := range inscricoes {
        relatorio[i] = gin.H{
            "UsuarioID": inscricao.UsuarioID,
            "Nome":      findUsuarioNameByID(usuarios, inscricao.UsuarioID),
            "Status":    inscricao.Status,
            "Data":      inscricao.Data,
            "Hora":      inscricao.Hora,
        }
    }

    c.JSON(http.StatusOK, relatorio)
}

// Função auxiliar para encontrar o nome do usuário pelo ID
func findUsuarioNameByID(usuarios []models.Usuario, usuarioID uint) string {
    for _, usuario := range usuarios {
        if usuario.ID == usuarioID {
            return usuario.Nome
        }
    }
    return "Nome não encontrado"
}
