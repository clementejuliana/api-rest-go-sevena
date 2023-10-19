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

	tipoUsuario := r.Group("/tipoUsuario")
	{
		tipoUsuario.GET("/", controllers.ExibirTipoUsuario)
		tipoUsuario.POST("/", controllers.CriarTipoUsuario)
		tipoUsuario.GET("/:id", controllers.BuscarTipoUsuarioPorID)
		tipoUsuario.DELETE("/:id", controllers.DeleteTipoUsuario)
		tipoUsuario.PATCH("/:id", controllers.EditarTipoUsuario)
	}

	locals := r.Group("/local")
	{
		locals.GET("/", controllers.ExibirLocal)
		locals.POST("/",controllers.CriarNovoLocal)
		locals.GET("/:id",controllers.BuscarLocalPorID)
		locals.DELETE("/:id",controllers.DeleteLocal)
		locals.PATCH("/:id",controllers.EditarLocal)
	}

	cidades := r.Group("/cidade")
	{
		cidades.GET("/",controllers.ExibirCidade)
		cidades.POST("/",controllers.CriarCidade)
		cidades.GET("/:id",controllers.BuscarCidadePorID)
		cidades.DELETE("/:id",controllers.DeleteCidade)
		cidades.PATCH("/:id",controllers.EditarCidade)
	}

	eventos := r.Group("/evento")
	{
		eventos.GET("/",controllers.ExibirEventos)
		eventos.POST("/",controllers.CriarNovoEvento)
		eventos.GET("/:id",controllers.BuscarEventoPorID)
		eventos.DELETE("/:id",controllers.DeleteEvento)
		eventos.PATCH("/:id",controllers.EditarEvento)
	}



	r.Run()

}
