package models

import "gorm.io/gorm"

type JobFieldValue struct {
	gorm.Model
	JobFieldID uint
	Value      string

	ApplicationID uint
}
