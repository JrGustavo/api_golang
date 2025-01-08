package main

import (
	"github.com/JrGustavo/api_golang/data"
	"github.com/JrGustavo/api_golang/models"
	"github.com/JrGustavo/api_golang/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env file")
	}

	data.ConectarPostgres()

	data.DB.AutoMigrate(&models.Rol{})
	data.DB.AutoMigrate(&models.Usuario{})

	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", rutas))

}
