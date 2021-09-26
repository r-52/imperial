package models

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	Name string

	Uid string

	ApplicationID uint

	OpenTill sql.NullTime
	IsClosed bool

	TimeChosenAt         sql.NullTime
	CloseAfterTimePicked bool

	NotifyWhenDoneEmail string
	CompanyPersonID     uint
	ScheduleEntry       []ScheduleEntry
}

func (c *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
