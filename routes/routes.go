package routes

import (
	"github.com/clementejuliana/api-rest-go-sevena/controllers"
	"github.com/gin-gonic/gin"
)

// função que lida com as requisões
func HandleRequests()  {
r := gin.Default()
r.GET("/usuarios", controllers.ExibirUsuario)
r.GET("/:nome", controllers.Saudacao)
r.POST("/usuarios", controllers.CriarNovoAluno)
r.GET("/usuarios/:id", controllers.BuscarUsuarioPorID)
r.DELETE("/usuarios/:id", controllers.DeleteUsuario)
r.PATCH("/usuarios/:id", controllers.EditarUsuario)
r.GET("usuarios/cpf/:cpf", controllers.BuscarUsuarioPorCPF)
r.Run()

}