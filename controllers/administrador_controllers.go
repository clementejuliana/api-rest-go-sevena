package controllers

import (
	"net/http"

	"github.com/clementejuliana/api-rest-go-sevena/authentication"
	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/gin-gonic/gin"
)

func ExibirAdministrador(c *gin.Context) {
	var administrador []models.Administrador
	databasee.DB.Find(&administrador)
	c.JSON(200, administrador)
}

// criar esse novo administrador
func CriarNovoAdministrador(c *gin.Context) {
	var administrador models.Administrador
	if err := c.ShouldBindJSON(&administrador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Create(&administrador)
	c.JSON(http.StatusOK, administrador)
}

func BuscarAdministradorPorID(c *gin.Context) {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.First(&administrador, id)

	if administrador.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Administrador Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, administrador)
}

func DeleteAdministrador(c *gin.Context) {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.Delete(&administrador, id)
	c.JSON(http.StatusOK, gin.H{"data": "Administrador deletado com sucesso"})
}

func EditarAdministrador(c *gin.Context)  {
	var administrador models.Administrador
	id := c.Params.ByName("id")
	databasee.DB.First(&administrador, id)

	if err := c.ShouldBindJSON(&administrador); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
    databasee.DB.Model(&administrador).UpdateColumns(administrador)
	c.JSON(http.StatusOK, administrador)
	
}

func ExibirUsuariosAdmin(c *gin.Context) {
	var usuarios []models.Usuario
	databasee.DB.Find(&usuarios)
	c.JSON(http.StatusOK, usuarios)
}

func DeletarUsuarioAdmin(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Usuário deletado com sucesso"})
}


func CriarUsuarioAdmin(c *gin.Context) {
	var novoUsuario models.Usuario
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Defina o papel do usuário como "admin"
	novoUsuario.Papel = "admin"

	if err := novoUsuario.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Resto da lógica para criar o usuário...
}

func AdicionarAdministrador(c *gin.Context) {
	// Verifique se o usuário atual tem permissão para adicionar administradores
	if !usuarioEhAdminAtual(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Você não tem permissão para adicionar administradores"})
		return
	}

	// Parseie os dados do novo administrador do corpo da solicitação
	var novoAdmin models.Usuario
	if err := c.ShouldBindJSON(&novoAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Defina o papel do novo administrador como "admin"
	novoAdmin.Papel = "admin"

	// Execute a lógica para adicionar o novo administrador ao banco de dados
	if err := databasee.DB.Create(&novoAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar administrador"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Administrador adicionado com sucesso"})
}

// Função auxiliar para verificar se o usuário atual é um administrador
func usuarioEhAdminAtual(c *gin.Context) bool {
	// Obtenha o ID do usuário a partir do token
	usuarioID, err := authentication.ExtrairUsuarioID(c)
	if err != nil {
		return false
	}

	// Consulte o banco de dados para obter informações sobre o usuário
	var usuario models.Usuario
	if err := databasee.DB.First(&usuario, usuarioID).Error; err != nil {
		return false
	}

	// Verifique se o usuário tem o papel de administrador
	return usuario.Papel == "admin"
}