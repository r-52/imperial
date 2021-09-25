package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	Name string

	ApplicationID uint

	OpenTill sql.NullTime
	IsClosed bool

	TimeChosenAt         sql.NullTime
	CloseAfterTimePicked bool

	NotifyWhenDoneEmail string
	CompanyPersonID     uint
	ScheduleEntry       []ScheduleEntry
}
