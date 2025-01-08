package routes

import (
	"github.com/JrGustavo/api_golang/controllers"
	"github.com/JrGustavo/api_golang/middleware"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()

	api.HandleFunc("", controllers.InitRoute).Methods("GET")

	apiRoles := api.PathPrefix("/roles").Subrouter()
	apiRoles.HandleFunc("", middleware.SetMiddlewareJSONAuthentication(controllers.GetRoles)).Methods("GET")
	apiRoles.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.GetRol)).Methods("GET")
	apiRoles.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.UpdateRol)).Methods("PUT")
	apiRoles.HandleFunc("", middleware.SetMiddlewareJSONAuthentication(controllers.NewRol)).Methods("POST")
	apiRoles.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.DeleteRol)).Methods("DELETE")

	apiUsuarios := api.PathPrefix("/usuarios").Subrouter()
	apiUsuarios.HandleFunc("", middleware.SetMiddlewareJSONAuthentication(controllers.GetUsuarios)).Methods("GET")
	apiUsuarios.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.GetUsuario)).Methods("GET")
	apiUsuarios.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.UpdateUsuario)).Methods("PUT")
	apiUsuarios.HandleFunc("", middleware.SetMiddlewareJSONAuthentication(controllers.NewUsuario)).Methods("POST")
	apiUsuarios.HandleFunc("/{id}", middleware.SetMiddlewareJSONAuthentication(controllers.DeleteUsuario)).Methods("DELETE")

	apiAuth := api.PathPrefix("/auth").Subrouter()

	apiAuth.HandleFunc("/login", controllers.Login).Methods("POST")
	apiAuth.HandleFunc("/register", controllers.NewUsuario).Methods("POST")

	return rutas

}
