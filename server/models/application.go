package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model

	AppliedAt                      sql.NullTime
	ConfirmationSentAt             sql.NullTime
	ApplicationOpenedAt            sql.NullTime
	ApplicationFirstResponseSentAt sql.NullTime
	ApplicationRejectedAt          sql.NullTime

	FirstName string
	LastName  string
	Email     string

	AcceptedByPersonID uint
	RejectedByPersonID uint
	InvitedByPersonID  uint

	JobFieldValue []JobFieldValue
	JobID         uint

	InternalNote       sql.NullString
	InternalStarReview sql.NullInt16

	ApplicationStatusID uint

	Schedule []Schedule
	Tag      []Tag
}
