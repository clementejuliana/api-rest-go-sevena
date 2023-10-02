package main

import (
	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/routes"
)

func main() {
	databasee.ConexaoBD()
	routes.HandleRequests()

}
