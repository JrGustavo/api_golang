package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var CONNECTION_STRING = "host=localhost user=postgres password=1234 dbname=api_golang port=5433 sslmode=disable TimeZone=America/Bogota"
var DB *gorm.DB

func ConectarPostgres() {
	var error error
	DB, error = gorm.Open(postgres.Open(CONNECTION_STRING), &gorm.Config{})
	if error != nil {
		log.Fatal("Error al conectar con la base de datos")
	} else {
		log.Println("Conectado a la base de datos")
	}
}
