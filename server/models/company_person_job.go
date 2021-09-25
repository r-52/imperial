package models

import "gorm.io/gorm"

type CompanyPersonJob struct {
	gorm.Model

	CompanyPersonID uint
	JobID           uint
}
