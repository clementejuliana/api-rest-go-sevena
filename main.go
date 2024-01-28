package main

import (
	"fmt"
	"os"


	"github.com/clementejuliana/api-rest-go-sevena/databasee"
	"github.com/clementejuliana/api-rest-go-sevena/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	databasee.ConexaoBD()
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

	// Criar o roteador
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Register routes
	r.GET("/", func(c *gin.Context) {
		// Set the response content type
		c.Header("Content-type", "application/json")

		// Respond with a JSON object
		c.JSON(200, gin.H{"message": "Hello, world!"})
	})

	//databasee.ConexaoBD()
	routes.HandleRequests()

	// Configurar o modo "release" do Gin
	gin.SetMode(gin.ReleaseMode)

	// Rota padrão
	r.GET("/", func(c *gin.Context) {
		fmt.Println("O servidor está rodando.")
	})

	// Iniciar o servidor
	r.Run()
}
