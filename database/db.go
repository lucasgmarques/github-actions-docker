package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  dbname := os.Getenv("DB_NAME")
  port := os.Getenv("DB_PORT")

	stringDeConexao := "host="+ host +" user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})

}
