package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type JobField struct {
	gorm.Model

	JobID   uint
	FieldID uint

	Position uint

	Label       string
	Payload     sql.NullString // for select stuff
	Placeholder sql.NullString
	HelpText    sql.NullString
	IsRequired  bool

	JobFieldValue []JobFieldValue
}
