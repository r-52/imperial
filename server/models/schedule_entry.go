package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ScheduleEntry struct {
	gorm.Model
	Date       time.Time      `json:"date"`
	IsChosen   bool           `json:"isChosen"`
	Note       sql.NullString `json:"note"`
	ScheduleID uint           `json:"-"`
}
