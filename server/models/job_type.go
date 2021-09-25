package models

import "gorm.io/gorm"

type JobType struct {
	gorm.Model

	Name string

	CompanyID uint
}
