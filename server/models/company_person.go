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

	CompanyID                uint
	CompanyPersonJob         []CompanyPersonJob
	AcceptedApplication      []Application `gorm:"foreignKey:AcceptedByPersonID"`
	RejectedApplication      []Application `gorm:"foreignKey:RejectedByPersonID"`
	InvitedApplication       []Application `gorm:"foreignKey:InvitedByPersonID"`
	Schedule                 []Schedule
	ApplicationInterviewNote []ApplicationInterviewNote
}
