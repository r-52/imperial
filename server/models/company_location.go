package models

import "gorm.io/gorm"

type CompanyLocation struct {
	gorm.Model

	CompanyID uint

	Name    string
	Street1 string
	Street2 string
	Zip     string
	City    string
	County  string
	Country string

	IsMainLocation bool

	JobCompanyLocation []JobCompanyLocation
}
