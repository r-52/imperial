package models

import "gorm.io/gorm"

type ApplicationStatus struct {
	gorm.Model
	Name string `json:"name"`

	Application []Application `json:"-"`
}
