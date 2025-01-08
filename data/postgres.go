package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConectarPostgres() {

	var error error
	DB, error = gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if error != nil {
		log.Fatal("Error al conectar con la base de datos")
	} else {
		log.Println("Conectado a la base de datos")
	}
}
