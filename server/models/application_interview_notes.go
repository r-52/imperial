package models

import "gorm.io/gorm"

type ApplicationInterviewNote struct {
	gorm.Model
	CompanyPersonID *uint  `json:"person"`
	InterviewTypeID *uint  `json:"interViewType"`
	Note            string `json:"note"`
}
