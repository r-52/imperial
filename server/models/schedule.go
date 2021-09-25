package models

import "database/sql"

type Schedule struct {
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
