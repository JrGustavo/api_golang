package main

import (
	"github.com/JrGustavo/api_golang/data"
	"github.com/JrGustavo/api_golang/models"
	"github.com/JrGustavo/api_golang/routes"
	"log"
	"net/http"
)

func main() {
	data.ConectarPostgres()

	data.DB.AutoMigrate(&models.Rol{})
	data.DB.AutoMigrate(&models.Usuario{})

	rutas := routes.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", rutas))

}
