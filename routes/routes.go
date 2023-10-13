package routes

import (
	"github.com/clementejuliana/api-rest-go-sevena/controllers"
	"github.com/gin-gonic/gin"
)

// função que lida com as requisões
func HandleRequests() {
	r := gin.Default()
	// Grupo de rotas de usuário
	user := r.Group("/usuario")
	{
		user.GET("/", controllers.ExibirUsuario)
		user.GET("/nome/:nome", controllers.Saudacao)
		user.POST("/", controllers.CriarNovoUsuario)
		user.GET("/:id", controllers.BuscarUsuarioPorID)
		user.DELETE("/:id", controllers.DeleteUsuario)
		user.PATCH("/:id", controllers.EditarUsuario)
		user.GET("usuario/cpf/:cpf", controllers.BuscarUsuarioPorCPF)
	}

	// Grupo de rotas de estado
	estado := r.Group("/estado")
	{
		estado.GET("/", controllers.ExibirEstado)
		estado.POST("/", controllers.CriarNovoEstado)
		estado.GET("/:id", controllers.BuscarEstadoPorID)
		estado.DELETE("/:id", controllers.DeleteEstado)
		estado.PATCH("/:id", controllers.EditarEstado)
	}

	// Grupo de rotas de instituição
	instituicao := r.Group("/instituicao")
	{
		instituicao.GET("/", controllers.ExibirInstuicao)
		instituicao.POST("/", controllers.CriarNovaInstituicao)
		instituicao.GET("/:id", controllers.BuscarInstuicaoPorID)
		instituicao.DELETE("/:id", controllers.DeleteInstituicao)
		instituicao.PATCH("/:id", controllers.EditarInstituicao)
	}

	r.Run()

}
