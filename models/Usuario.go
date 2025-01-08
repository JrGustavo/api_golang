package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type Usuario struct {
	gorm.Model
	Nombre   string `json:"nombre" gorm:"size:100;not null"`
	Correo   string `json:"correo" gorm:"size:100;unique;not null"`
	Password string `json:"password" gorm:"not null"` // Eliminar el default
	RolId    uint   `json:"rol_id"`
	Rol      Rol    `json:"rol"`
}

type UsuarioResponse struct {
	ID        uint      `json:"id"`
	Nombre    string    `json:"nombre"`
	Correo    string    `json:"correo"`
	RolId     uint      `json:"rol_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Rol       Rol       `json:"rol"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

// Hash genera un hash del password utilizando bcrypt.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerificarPassword compara el hash almacenado con el password plano.
func VerificarPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}

// BeforeSave se ejecuta automáticamente antes de guardar un usuario en la base de datos.
func (u *Usuario) BeforeSave(tx *gorm.DB) error {
	// Hashear la contraseña antes de guardarla
	passwordHashed, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(passwordHashed)
	return nil
}

// Prepare limpia y normaliza los campos del modelo Usuario.
func (u *Usuario) Prepare() {
	u.Nombre = html.EscapeString(strings.ToUpper(strings.TrimSpace(u.Nombre)))
	u.Correo = html.EscapeString(strings.TrimSpace(u.Correo))
}
