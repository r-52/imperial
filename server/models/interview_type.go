package models

import "gorm.io/gorm"

type InterviewType struct {
	gorm.Model
	Name string

	CompanyID                uint
	ApplicationInterviewNote []ApplicationInterviewNote
}
