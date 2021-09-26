package models

import "gorm.io/gorm"

type JobCompanyLocation struct {
	gorm.Model
	CompanyLocationID uint
	CompanyLocation   CompanyLocation

	JobID uint
	Job   *Job
}
