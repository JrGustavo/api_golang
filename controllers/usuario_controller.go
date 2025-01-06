package controllers

import (
	"encoding/json"
	"github.com/JrGustavo/api_golang/data"
	"github.com/JrGustavo/api_golang/models"
	"github.com/JrGustavo/api_golang/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuarios []models.Usuario
	data.DB.Preload("Rol").Find(&usuarios)
	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Lista de usuarios",
		StatusCode: http.StatusOK,
		Data:       usuarios,
	})
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var usuario []models.Usuario
	data.DB.Preload("Rol").Find(&usuario, params["id"])
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data.DB.Model(&usuario).Association("Roles").Find(&usuario.RolId)

	json.NewEncoder(w).Encode(&usuario)

}

func NewUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	createdUsuario := data.DB.Create(&usuario)
	err := createdUsuario.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

	}
	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Usuario creado con exito",
		StatusCode: http.StatusOK,
		Data:       usuario,
	})
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al decodificar el usuario",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		})
		return
	}

	var usuarioExistente models.Usuario
	if err := data.DB.Preload("Rol").First(&usuarioExistente, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Usuario no encontrado",
			StatusCode: http.StatusNotFound,
			Data:       err.Error(),
		})
		return
	}

	usuarioExistente.Nombre = usuario.Nombre
	usuarioExistente.Correo = usuario.Correo
	usuarioExistente.RolId = usuario.RolId
	usuarioExistente.Password = usuario.Password
	if err := data.DB.Save(&usuarioExistente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al actualizar",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		})
		return
	}

	respuesta := utils.Respuesta{
		Msg:        "Usuario actualizado con éxito",
		StatusCode: http.StatusOK,
		Data:       usuarioExistente,
	}
	json.NewEncoder(w).Encode(&respuesta)

}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var usuario models.Usuario
	data.DB.Preload("Rol").First(&usuario, params["id"])

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data.DB.Delete(&usuario)
	w.WriteHeader(http.StatusOK)

	}


}
