package models

import "gorm.io/gorm"

type Knowledge struct {
	gorm.Model

	Name      string
	CompanyID uint
}
