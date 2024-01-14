package controllers

import (
	"net/http"
	"strconv"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirEventos(c *gin.Context) {
	var evento []models.Evento
	databasee.DB.Find(&evento)
	c.JSON(200, evento)

}

//exibir uma mensagem quando está passando um valoe não valido
func SaudacaoEventos(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Tudo bem " + nome + ", tudo beleza?",
	})
}

// criar esse novo aluno
func CriarNovoEvento(c *gin.Context) {
	var evento models.Evento
	if err := c.ShouldBindJSON(&evento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := evento.Preparar(databasee.DB); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&evento)
	c.JSON(http.StatusOK, evento)
}

func BuscarEventoPorID(c *gin.Context) {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.First(&evento, id)

	if evento.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Evento Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, evento)
}

func DeleteEvento(c *gin.Context) {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.Delete(&evento, id)
	c.JSON(http.StatusOK, gin.H{"data": "Evento deletado com sucesso"})
}

func EditarEvento(c *gin.Context)  {
	var evento models.Evento
	id := c.Params.ByName("id")
	databasee.DB.First(&evento, id)

	if err := c.ShouldBindJSON(&evento); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&evento).UpdateColumns(evento)
	c.JSON(http.StatusOK, evento)
	
}


// ObterUsuariosInscritosNoEvento obtém a lista de usuários inscritos em um evento
func ObterUsuariosInscritosNoEventoo(c *gin.Context) {
	eventoID, err := strconv.ParseUint(c.Param("evento_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	// Consulta ao banco de dados para obter as inscrições no evento
	var inscricoes []models.InscricaoEmEvento
	if err := databasee.DB.Where("evento_id = ?", eventoID).Find(&inscricoes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter inscrições no evento"})
		return
	}

	// Lista para armazenar informações dos usuários
	var usuarios []gin.H

	// Itera sobre as inscrições para obter informações dos usuários
	for _, inscricao := range inscricoes {
		var usuario models.Usuario
		if err := databasee.DB.First(&usuario, inscricao.UsuarioID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter informações do usuário"})
			return
		}

		// Adiciona informações do usuário à lista
		usuarios = append(usuarios, gin.H{
			"id":    usuario.ID,
			"nome":  usuario.Nome,
			"email": usuario.Email,
			// Adicione outros campos conforme necessário
		})
	}

	c.JSON(http.StatusOK, gin.H{"usuarios_inscritos": usuarios})
}


func ObterUsuariosInscritosNoEvento(c *gin.Context) {
    eventoID, err := strconv.ParseUint(c.Param("evento_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
        return
    }

    // Consulta ao banco de dados para obter as inscrições no evento
    rows, err := databasee.DB.Raw("SELECT usuario.id, usuario.nome, usuario.email FROM inscricoes_em_evento JOIN usuarios ON inscricoes_em_evento.usuario_id = usuarios.id WHERE inscricoes_em_evento.evento_id = ?", eventoID).Rows()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter inscrições no evento"})
        return
    }
    defer rows.Close()

    // Itera sobre os registros da consulta
    var usuarios []gin.H
    for rows.Next() {
        var id uint64
        var nome string
        var email string
        if err := rows.Scan(&id, &nome, &email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao iterar sobre registros da consulta"})
            return
        }
        usuario := gin.H{
            "id": id,
            "nome": nome,
            "email": email,
        }
        usuarios = append(usuarios, usuario)
    }

    c.JSON(http.StatusOK, gin.H{"usuarios_inscritos": usuarios})
}


// // Função auxiliar para encontrar o nome do usuário pelo ID
// func findUsuarioNameByID(usuarios []models.Usuario, usuarioID uint) string {
//     for _, usuario := range usuarios {
//         if usuario.ID == usuarioID {
//             return usuario.Nome
//         }
//     }
//     return "Nome não encontrado"
// }

// ObterUsuariosInscritosNoEvento obtém a lista de usuários inscritos em um evento

