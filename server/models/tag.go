package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model

	Name  string `json:"name"`
	Color string `json:"color"`

	ApplicationID *uint `json:"-"`
	JobID         *uint `json:"-"`
}
