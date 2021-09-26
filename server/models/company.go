package models

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Name                   string
	DataProtectionTemplate sql.NullString
	TermsOfServiceTemplate sql.NullString
	Imprint                sql.NullString
	Phone                  sql.NullString
	Email                  sql.NullString
	AdminEmail             string
	Slogan                 sql.NullString

	Uid string

	CompanyUrl   string
	FacebookUrl  string
	TwitterUrl   string
	InstagramUrl string

	CustomStyleSheetInjection sql.NullString
	CustomLogoUrl             sql.NullString

	CompanyLocation []CompanyLocation
	ContactPerson   []CompanyPerson
	JobType         []JobType
	Job             []Job
	MailTemplate    []MailTemplate
	InterviewType   []InterviewType
	Knowledge       []Knowledge
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
