package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type MailTemplate struct {
	gorm.Model
	Name             string
	Description      sql.NullString
	InternalUsageKey uint
	Subject          string
	Body             string
	Bcc              string
	Cc               string

	CompanyID uint
}
