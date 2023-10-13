package main

import (
	"fmt"
	"os"

	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregue as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente:", err)
	}

	// Obtenha o valor da variável de ambiente GIN_MODE
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "" {
		ginMode = "debug" // Valor padrão se não estiver definido no .env
	}

	// Configurar o modo Gin
	gin.SetMode(ginMode)

	databasee.ConexaoBD()
	routes.HandleRequests()

	// Configurar o modo "release" do Gin
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("O servidor está rodando.")
	})

	// Start the server
	r.Run()

}

