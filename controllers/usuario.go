package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/services"
	"github.com/gin-gonic/gin"
)

func ExibirUsuario(c *gin.Context) {
	var usuario []models.Usuario
	databasee.DB.Find(&usuario)
	c.JSON(200, usuario)

}

// criar esse novo aluno
func CriarNovoUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := usuario.Preparar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var usuarioExistente models.Usuario
	email := databasee.DB.Where("email = ?", usuario.Email).Find(&usuarioExistente)
	if email.RowsAffected > 0 {
		c.JSON(400, gin.H{"MENSAGEM": "Esse email já está cadastrado no sistema!"})
		return
	}

	usuario.Senha = services.SHA256Encoder(usuario.Senha)
	databasee.DB.Create(&usuario)
	c.JSON(http.StatusOK, gin.H{
		"data": "Cadastro realizado com sucesso!"})
}

func BuscarUsuarioPorID(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if usuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func DeleteUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario deletado com sucesso"})
}

func EditarUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	databasee.DB.Model(&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)

}

func BuscarUsuarioPorCPF(c *gin.Context) {
	var usuario models.Usuario
	cpf := c.Param("cpf")
	databasee.DB.Where(&models.Usuario{CPF: cpf}).First(&usuario)

	if usuario.ID == 00 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Usuario Não encontrando"})
		return
	}
	c.JSON(http.StatusOK, usuario)

}

func IniciaRecuperacaoSenha(c *gin.Context) {
	var usuario models.Usuario
	id := c.Params.ByName("id")
	databasee.DB.First(&usuario, id)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// Gera um token único para recuperação de senha
	tokenRecuperacao, err := services.GeraTokenRecuperacaoSenha()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token de recuperação de senha"})
		return
	}

	// Define a data de expiração para 1 hora a partir de agora
	dataExpiracao := time.Now().Add(time.Hour * 1)

	// Salva as informações de recuperação de senha no banco de dados
	recuperacaoSenha := models.RecuperacaoSenha{
		UsuarioID:     int(usuario.ID),
		Token:         tokenRecuperacao,
		Email:         usuario.Email,
		DataExpiracao: dataExpiracao,
	}

	// Salva no banco de dados usando o GORM
	if err := databasee.DB.Create(&recuperacaoSenha).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar informações de recuperação de senha"})
		return
	}

	// Gera o link de recuperação
	linkRecuperacao := fmt.Sprintf("http://localhost:8080/usuario/inicia-recuperacao-senha?token=%s", tokenRecuperacao)

	// Envia um e-mail ao usuário com o link de recuperação
	if err := services.EnviaEmailRecuperacaoSenha(usuario.Email, linkRecuperacao); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar e-mail de recuperação de senha"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Processo de recuperação de senha iniciado com sucesso"})
}



