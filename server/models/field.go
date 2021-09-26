package models

import "gorm.io/gorm"

type Field struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
