package models

import "gorm.io/gorm"

type ApplicationInterviewNote struct {
	gorm.Model
	CompanyPersonID uint
	InterviewTypeID uint
}
