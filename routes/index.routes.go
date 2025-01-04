package routes

import (
	"github.com/JrGustavo/api_golang/controllers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("", controllers.InitRoute).Methods("GET")

	apiRoles := api.PathPrefix("/roles").Subrouter()
	apiRoles.HandleFunc("", controllers.GetRoles).Methods("GET")
	apiRoles.HandleFunc("/{id}", controllers.GetRol).Methods("GET")
	apiRoles.HandleFunc("", controllers.NewRol).Methods("POST")
	apiRoles.HandleFunc("/{id}", controllers.DeleteRol).Methods("DELETE")

	return rutas

}
