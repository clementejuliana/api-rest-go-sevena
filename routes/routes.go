package routes

import (
	"github.com/clementejuliana/api-rest-go-sevena/controllers"
	"github.com/gin-gonic/gin"
)

// função que lida com as requisões
func HandleRequests() {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	// Grupo de rotas de usuário
	user := r.Group("/usuario")
	{
		user.GET("/", controllers.ExibirUsuario)
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
		estado.GET("/carrega-estado", controllers.CarregarEstados)
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
		locals.POST("/", controllers.CriarNovoLocal)
		locals.GET("/:id", controllers.BuscarLocalPorID)
		locals.DELETE("/:id", controllers.DeleteLocal)
		locals.PATCH("/:id", controllers.EditarLocal)
	}

	cidades := r.Group("/cidade")
	{
		cidades.GET("/", controllers.ExibirCidade)
		cidades.POST("/", controllers.CriarCidade)
		cidades.GET("/:id", controllers.BuscarCidadePorID)
		cidades.DELETE("/:id", controllers.DeleteCidade)
		cidades.PATCH("/:id", controllers.EditarCidade)
		cidades.POST("/carrega-cidades", controllers.CarregarCidades)
	}

	eventos := r.Group("/evento")
	{
		eventos.GET("/", controllers.ExibirEventos)
		eventos.POST("/", controllers.CriarNovoEvento)
		eventos.GET("/:id", controllers.BuscarEventoPorID)
		eventos.DELETE("/:id", controllers.DeleteEvento)
		eventos.PATCH("/:id", controllers.EditarEvento)
	}

	administradores := r.Group("/administrador")
	{
		administradores.GET("/", controllers.ExibirAdministrador)
		administradores.POST("/", controllers.CriarNovoAdministrador)
		administradores.GET("/:id", controllers.BuscarAdministradorPorID)
		administradores.DELETE("/:id", controllers.DeleteAdministrador)
		administradores.PATCH("/:id", controllers.EditarAdministrador)
	}

	atividades := r.Group("/atividade")
	{
		atividades.GET("/", controllers.ExibirAtividade)
		atividades.POST("/", controllers.CriarNovaAtividade)
		atividades.GET("/:id", controllers.BuscarAtividadePorID)
		atividades.DELETE("/:id", controllers.DeleteAtividade)
		atividades.PATCH("/:id", controllers.EditarAtividade)
	}

controlePresencas := r.Group("/controlePresenca")
{
	controlePresencas.GET("/",controllers.ExibirControlePresenca)
	controlePresencas.POST("/",controllers.CriarNovoControlePresenca)
	controlePresencas.GET("/:id",controllers.BuscarControlePresencaPorID)
	controlePresencas.DELETE("/:id",controllers.DeleteControlePresenca)
	controlePresencas.PATCH("/:id",controllers.EditarControlePresenca)
}

inscricaoEmEventos := r.Group("/inscricaoEmEventos")
{
	inscricaoEmEventos.GET("/",controllers.ExibirInscricaoEmEventos)
	inscricaoEmEventos.POST("/",controllers.CriarNovaInscricaoEmEventos)
	inscricaoEmEventos.GET("/:id",controllers.BuscarInscricaoEmEventosPorID)
	inscricaoEmEventos.DELETE("/:id",controllers.DeleteInscricaoEmEventos)
	inscricaoEmEventos.PATCH("/:id",controllers.EditarInscricaoEmEventos)
}

inscricaoEmAtividades := r.Group("/inscricaoEmAtividades")
{
	inscricaoEmAtividades.GET("/",controllers.ExibirInscricaoEmAtividade)
	inscricaoEmAtividades.POST("/",controllers.CriarNovaInscricaoEmAtividade)
	inscricaoEmAtividades.GET("/:id",controllers.BuscarInscricaoEmAtividadePorID)
	inscricaoEmAtividades.DELETE("/:id",controllers.DeleteInscricaoEmAtividade)
	inscricaoEmAtividades.PATCH("/:id",controllers.EditarInscricaoEmAtividade)
}
tipoAtividades := r.Group("/tipoAtividades")
{
	tipoAtividades.GET("/",controllers.ExibirTipoAtividade)
	tipoAtividades.POST("/",controllers.CriarNovoTipoAtividade)
	tipoAtividades.GET("/:id",controllers.BuscarTipoAtividadePorID)
	tipoAtividades.DELETE("/:id",controllers.DeleteTipoAtividade)
	tipoAtividades.PATCH("/:id",controllers.EditarTipoAtividade)
}


	r.Run()

}
