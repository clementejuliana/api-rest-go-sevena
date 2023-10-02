package main

import (
	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/models"
	"github.com/clementejuliana/api-rest-go-sevena/routes"
)

func main() {
	databasee.ConexaoBD()
	models.Usuarios = []models.Usuario{
		{
			Status:       "ON",
			Nome:         "Thiago Gouveia",
			CPF:          "0000000000000000",
			RG:           "470000000000",
			Genero:       "M",
			Email:        "THIAGO@GMAIL.COM",
			Telefone:     "4498883406",
			Escolaridade: "ensino medio",
			Profissao:    "dev",
			FotoPerfil:   "ok",
		},
		{
			Status:       "ON",
			Nome:         "Juliana",
			CPF:          "0000000000000000",
			RG:           "470000000000",
			Genero:       "M",
			Email:        "THIAGO@GMAIL.COM",
			Telefone:     "4498883406",
			Escolaridade: "ensino medio",
			Profissao:    "dev",
			FotoPerfil:   "ok",
		},
	}
	routes.HandleRequests()

}
