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
r.POST("/usuarios", controllers.CriarNovoUsuario)
r.GET("/usuarios/:id", controllers.BuscarUsuarioPorID)
r.DELETE("/usuarios/:id", controllers.DeleteUsuario)
r.PATCH("/usuarios/:id", controllers.EditarUsuario)
r.GET("usuarios/cpf/:cpf", controllers.BuscarUsuarioPorCPF)

r.GET("/estados", controllers.ExibirEstado)
r.GET("/:nome", controllers.SaudacaoEstado)
r.POST("/estados", controllers.CriarNovoEstado)
r.GET("/estados", controllers.BuscarEstadoPorID)
r.DELETE("/estados/:id", controllers.DeleteEstado)
r.PATCH("/estados/:id", controllers.EditarEstado)


r.GET("/instituicao", controllers.ExibirInstuicao)
r.GET("/:nome", controllers.SaudacaoInstituicao)
r.POST("/instituicao", controllers.CriarNovaInstituicao)
r.GET("/instituicao", controllers.BuscarInstuicaoPorID)
r.DELETE("/instituicao/:id", controllers.DeleteInstituicao)
r.PATCH("/instituicao/:id", controllers.EditarInstituicao)

r.Run()

}