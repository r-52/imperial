package models

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	Name string `json:"name"`

	Uid string `json:"uid" gorm:"uniqueIndex"`

	ApplicationID uint `json:"-"`

	OpenTill sql.NullTime `json:"openTill"`
	IsClosed bool         `json:"isClosed"`

	TimeChosenAt         sql.NullTime `json:"timeChosenAt"`
	CloseAfterTimePicked bool         `json:"closeAfterTimePicked"`

	NotifyWhenDoneEmail *string          `json:"notifyWhenDoneEmail"`
	CompanyPersonID     *uint            `json:"companyPerson"`
	ScheduleEntry       *[]ScheduleEntry `json:"entries"`
}

func (c *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
