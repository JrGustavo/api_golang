package controllers

import (
	"encoding/json"
	"github.com/JrGustavo/api_golang/utils"
	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Msg:        "Listado de roles",
		StatusCode: 200,
		Data:       "",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}

func NewRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Msg:        "Nuevo",
		StatusCode: 200,
		Data:       "Nuevo",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}

func GetRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Msg:        "Buscar",
		StatusCode: 200,
		Data:       "Buscar uno",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}

func DeleteRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Msg:        "Borrando",
		StatusCode: 200,
		Data:       "Eliminando ",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
