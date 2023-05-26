package database

import (
	"log"

	"github.com/Mariano-JR/sistema-de-login/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de Dados.")
	}

	DB.AutoMigrate(&models.User{})
}
