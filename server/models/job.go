package models

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model

	Slug string `json:"slug"`
	Uid  string `json:"uid" gorm:"uniqueIndex"`

	Title             string         `json:"title,omitempty"`
	Subline           sql.NullString `json:"subline,omitempty"`
	Intro             string         `json:"intro,omitempty"`
	Description       string         `json:"description,omitempty"`
	JobLogo           string         `json:"jobLogo,omitempty"`
	ApplyPossibleTill sql.NullTime   `json:"applyPossibleTill,omitempty"`

	ExternalJobUrlForApply sql.NullString `json:"externalJobUrlForApply,omitempty"`

	JobAvailableAt            *sql.NullTime `json:"jobAvailableAt,omitempty"`
	InternalJobAvailableAt    *sql.NullTime `json:"internalJobAvailableAt,omitempty"`
	IsAutoSwitchToPublicActiv bool          `json:"isAutoSwitchToPublicActiv"`

	IsJobPublic bool `json:"isJobPublic"`
	IsJobDraft  bool `json:"isJobDraft"`

	CompanyLocationID *uint               `json:"-"`
	CompanyID         *uint               `json:"-"`
	CompanyPersonJob  *[]CompanyPersonJob `json:"-"`
	KnowledgeID       *uint               `json:"-"`
	JobField          *[]JobField         `json:"-"`
	Tag               *[]Tag              `json:"-"`
}

func (c *Job) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
