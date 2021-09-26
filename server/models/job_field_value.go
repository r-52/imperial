package models

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type JobFieldValue struct {
	gorm.Model
	JobFieldID uint   `json:"field"`
	Value      string `json:"value"`

	Uid string `json:"uid" gorm:"uniqueIndex"`

	ApplicationID uint `json:"application"`
}

func (c *JobFieldValue) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
