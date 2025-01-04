package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type Usuario struct {
	gorm.Model

	ID       uint64 `json:"id" gorm:"primary_key;autoIncrement"`
	Nombre   string `json:"nombre" gorm:"size:100;not null"`
	Correo   string `json:"correo" gorm:"size:100;unique;not null"`
	Password string `json:"password" gorm:"default:true"`
	RolId    uint   `json:"rol_id"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerficarPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}

func (u *Usuario) BeforeSAve(tx *gorm.DB) error {
	passwordHashed, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(passwordHashed)
	return nil
}

func (u *Usuario) Prepare() error {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.ToUpper(strings.TrimSpace(u.Nombre)))
	u.Correo = html.EscapeString(strings.TrimSpace(u.Correo))
	return nil
}
