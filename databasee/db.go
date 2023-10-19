package databasee

import (
	"log"

	"github.com/clementejuliana/api-rest-go-sevena/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var(

	DB *gorm.DB
	err error

)

func ConexaoBD () {
	conexaoString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexaoString))
	if err != nil{
		log.Panic("erro ao conectar com o banco de dados") 
	}
	DB.AutoMigrate(&models.Usuario{})
	DB.AutoMigrate(&models.Estado{})
	DB.AutoMigrate(&models.Instituicao{})
	DB.AutoMigrate(&models.TipoUsuario{})//, &models.Atividade{}//)
	//DB.AutoMigrate(&models.Administrador{})
	//DB.AutoMigrate(&models.Atividade{})

	//DB.AutoMigrate(&models.Auth{})
	DB.AutoMigrate(&models.Cidade{})
	//DB.AutoMigrate(&models.ControlePresenca{})
	DB.AutoMigrate(&models.Evento{})
	//DB.AutoMigrate(&models.InscricaoEmAtividade{})
	//DB.AutoMigrate(&models.InscricaoEmEvento{})
	DB.AutoMigrate(&models.Local{})
	//DB.AutoMigrate(&models.Notificacao{})
	//DB.AutoMigrate(&models.RecuperacaoSenha{})
	//DB.AutoMigrate(&models.TipoAtividade{})
	
	
}



