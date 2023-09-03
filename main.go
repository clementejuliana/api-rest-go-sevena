package main

import (
	"fmt"

	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/routes"
)

func main() {

models.Notificacaos = []models.Notificacao{
	{NotificacaoID: 1 ,UsuarioID: 1, Conteudo: "SAI TCC, BORA PRODUDIZIR CARAMBA"},
	{NotificacaoID: 2 ,UsuarioID: 2, Conteudo: "SAI TCC, BORA PRODUDIZIR CARAMBA"},
}
r := routes.SetupRouter()
r.Run(":8000")
	fmt.Println("Iniciando o tcc sevena finalmente")

	//routes.HandleRequest()
}
