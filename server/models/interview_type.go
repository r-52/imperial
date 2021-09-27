package models

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type InterviewType struct {
	gorm.Model
	Name string

	Uid string

	CompanyID                uint
	ApplicationInterviewNote []ApplicationInterviewNote
}

func (c *InterviewType) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
