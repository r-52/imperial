package models

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Name                   string
	DataProtectionTemplate string
	TermsOfServiceTemplate string
	Imprint                string
	Phone                  string
	Email                  string
	AdminEmail             string
	Slogan                 string

	Uid string

	CompanyUrl   string
	FacebookUrl  string
	TwitterUrl   string
	InstagramUrl string

	CustomStyleSheetInjection string
	CustomLogoUrl             string

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
