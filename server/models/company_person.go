package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type CompanyPerson struct {
	gorm.Model

	FirstName    string
	LastName     string
	AvatarUrl    sql.NullString
	PasswordHash sql.NullString
	Email        string
	Phone        string
	Position     string

	CompanyID        uint
	CompanyPersonJob []CompanyPersonJob
	Application      []Application
	Schedule         []Schedule
}
