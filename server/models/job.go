package models

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model

	Slug string
	Uid  string

	Title             string
	Subline           sql.NullString
	Intro             string
	Description       string
	JobLogo           string
	ApplyPossibleTill sql.NullTime

	ExternalJobUrlForApply sql.NullString

	JobAvailableAt sql.NullTime

	CompanyLocationID uint
	CompanyID         uint
	CompanyPersonJob  []CompanyPersonJob
	KnowledgeID       uint
	JobField          []JobField
	Tag               []Tag
}

func (c *Job) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
