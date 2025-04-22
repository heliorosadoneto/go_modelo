package config

import (
	"git/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=123 dbname=teste port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco: ", err)
	}
	DB = db
	DB.AutoMigrate(models.User{}) // Cria a tabela se não existir
}
