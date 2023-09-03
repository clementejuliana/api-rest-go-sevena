package routes

import (
	
	

	"github.com/clementejuliana/api-rest-go-sevena/controllers"
	"github.com/gin-gonic/gin"
)

//func HandleRequest()  {
//	r := gin.Default()
//	http.HandleFunc("/", controllers.Home)
//	http.HandleFunc("/api/notificacaos", controllers.NotificacaoTodos)
//	log.Fatal(http.ListenAndServe(":8000", r))
//}


func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Configure as rotas usando o Gin
    
    r.GET("/", controllers.Home)
    r.GET("/api/notificacaos", controllers.NotificacaoTodos)
	r.GET("/api/notificacaos/{id}", controllers.NotificacaoTodos)
    return r
}