package models

import "gorm.io/gorm"

type Cita struct {
	gorm.Model
	Id        string `json:"id"`
	DNIType   string `json:"dnitype"`
	DNI       string `json:"dni"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}
