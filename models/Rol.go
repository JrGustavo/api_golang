package models

import "gorm.io/gorm"

// Rol representa los roles del sistema.
type Rol struct {
	gorm.Model
	Nombre   string    `json:"nombre" gorm:"unique;not null"` // Nombre único y no nulo
	Activo   bool      `json:"activo" gorm:"default:true"`    // Indica si el rol está activo
	Usuarios []Usuario `json:"usuarios"`                      // Relación con los usuarios
}

// TableName define el nombre de la tabla en la base de datos.
func (Rol) TableName() string {
	return "roles"
}
