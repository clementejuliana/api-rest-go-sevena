package controllers

import (
	
	"net/http"
	"strconv"
	"time"

	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    c.String(http.StatusOK, "HOME PAGE %s", time.Now())
}

func NotificacaoTodos(c *gin.Context) {
    c.JSON(http.StatusOK, models.Notificacaos)
}



func RetornarNotificacao(c *gin.Context) {
    // Obter o valor do parâmetro "id" da rota
    id := c.Param("id")
    // Retornar o ID da notificação
    c.String(http.StatusOK, id)

// Procurar a notificação com o ID especificado
	for _,notificacao :=range models.Notificacaos{
		if strconv.Itoa(notificacao.NotificacaoID) ==id{
			 // Retornar informações sobre a notificação encontrada
			c.JSON(http.StatusOK, notificacao)
			return
		}

	}
}
