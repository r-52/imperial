package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type JobField struct {
	gorm.Model

	JobID   *uint `json:"-"`
	FieldID *uint `json:"-"`

	Position uint `json:"position"`

	Label       string         `json:"label"`
	Payload     sql.NullString `json:"payload"` // for select stuff
	Placeholder sql.NullString `json:"placeholder"`
	HelpText    sql.NullString `json:"helpText"`
	IsRequired  bool           `json:"isRequired"`

	JobFieldValue *[]JobFieldValue `json:"-"`
}
