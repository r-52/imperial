package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type CompanyLocation struct {
	gorm.Model

	CompanyID uint

	Name    string
	Street1 string
	Street2 sql.NullString
	Zip     string
	City    string
	County  string
	Country string

	Job []Job
}
