package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model

	Title             string
	Subline           sql.NullString
	Intro             string
	Description       string
	JobLogo           string
	ApplyPossibleTill sql.NullTime

	ExternalJobUrlForApply sql.NullString

	JobAvailableAt sql.NullTime

	CompanyLocationID uint
	CompanyID         uint
	CompanyPersonJob  []CompanyPersonJob
	KnowledgeID       uint
	JobField          []JobField
	Tag               []Tag
}
