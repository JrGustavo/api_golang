package controllers

import (
	"encoding/json"
	"errors"
	"github.com/JrGustavo/api_golang/data"
	"github.com/JrGustavo/api_golang/models"
	"github.com/JrGustavo/api_golang/utils"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuarios []models.Usuario

	if err := data.DB.Preload("Rol").Find(&usuarios).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al obtener usuarios",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Lista de usuarios",
		StatusCode: http.StatusOK,
		Data:       usuarios,
	})
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var usuario models.Usuario

	if err := data.DB.Preload("Rol").First(&usuario, params["id"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.Respuesta{
				Msg:        "Usuario no encontrado",
				StatusCode: http.StatusNotFound,
				Data:       nil,
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.Respuesta{
				Msg:        "Error al obtener el usuario",
				StatusCode: http.StatusInternalServerError,
				Data:       err.Error(),
			})
		}
		return
	}

	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Usuario encontrado",
		StatusCode: http.StatusOK,
		Data:       usuario,
	})
}

func NewUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	// Verificar si el rol existe
	var rol models.Rol
	if err := data.DB.First(&rol, usuario.RolId).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Rol no encontrado",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		})
		return
	}

	// Crear el usuario
	if err := data.DB.Create(&usuario).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al crear el usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Usuario registrado con éxito",
		StatusCode: http.StatusOK,
		Data:       usuario,
	})
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var usuario models.Usuario
	var usuarioExistente models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al decodificar datos",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		})
		return
	}

	if err := data.DB.First(&usuarioExistente, params["id"]).Error; err != nil {
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
	if usuario.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
		usuarioExistente.Password = string(hashedPassword)
	}

	if err := data.DB.Save(&usuarioExistente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al actualizar el usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Usuario actualizado con éxito",
		StatusCode: http.StatusOK,
		Data:       usuarioExistente,
	})
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var usuario models.Usuario

	if err := data.DB.First(&usuario, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Usuario no encontrado",
			StatusCode: http.StatusNotFound,
			Data:       err.Error(),
		})
		return
	}

	if err := data.DB.Delete(&usuario).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Respuesta{
			Msg:        "Error al eliminar el usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(utils.Respuesta{
		Msg:        "Usuario eliminado con éxito",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}
