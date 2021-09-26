package models

import "gorm.io/gorm"

type JobType struct {
	gorm.Model

	Name string `json:"name"`

	CompanyID uint `json:"-"`
}
