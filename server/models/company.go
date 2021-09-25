package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Name                   string
	DataProtectionTemplate sql.NullString
	TermsOfServiceTemplate sql.NullString
	Imprint                sql.NullString
	Phone                  sql.NullString
	Email                  sql.NullString
	Slogan                 sql.NullString

	CustomStyleSheetInjection sql.NullString
	CustomLogoUrl             sql.NullString

	CompanyLocation []CompanyLocation
	ContactPerson   []CompanyPerson
	JobType         []JobType
	Job             []Job
}
