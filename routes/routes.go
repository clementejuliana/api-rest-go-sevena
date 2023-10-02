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
r.Run()

}