package job

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/r-52/imperial/models"
)

type JobViewModel struct {
	Uid                       string        `json:"uid,omitempty" validate:"required,min=2,max=128"`
	Slug                      string        `json:"slug,omitempty" validate:"required,min=2,max=1024"`
	Title                     string        `json:"title,omitempty" validate:"required,min=2"`
	Subline                   string        `json:"subline,omitempty"`
	Intro                     string        `json:"intro,omitempty"`
	Description               string        `json:"description,omitempty" validate:"required,min=2,max=8192"`
	JobLogo                   string        `json:"jobLogo,omitempty"`
	ApplyPossibleTill         sql.NullTime  `json:"applyPossibleTill,omitempty"`
	JobAvailableAt            *sql.NullTime `json:"jobAvailableAt,omitempty"`
	InternalJobAvailableAt    *sql.NullTime `json:"internalJobAvailableAt,omitempty"`
	IsJobPublic               bool          `json:"isJobPublic,omitempty"`
	IsJobDraft                bool          `json:"isJobDraft,omitempty"`
	IsAutoSwitchToPublicActiv bool          `json:"isAutoSwitchToPublicActiv,omitempty"`
	CompanyLocations          []string      `json:"locations,omitempty" validate:"required,dive,required"`
	Persons                   []string      `json:"persons,omitempty" validate:"required,dive,required"`
}

func (c *JobViewModel) Validate() (bool, error) {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *JobViewModel) ConvertToDbModel() models.Job {
	job := new(models.Job)
	job.ApplyPossibleTill = c.ApplyPossibleTill
	job.JobLogo = c.JobLogo
	job.Title = c.Title
	job.Subline = c.Subline
	job.Intro = c.Intro
	job.Description = c.Description
	job.InternalJobAvailableAt = c.InternalJobAvailableAt
	job.JobAvailableAt = c.JobAvailableAt
	job.IsJobPublic = c.IsJobPublic
	job.IsJobDraft = c.IsJobDraft
	job.IsAutoSwitchToPublicActiv = c.IsAutoSwitchToPublicActiv

	return *job
}
