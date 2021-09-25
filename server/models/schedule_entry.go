package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ScheduleEntry struct {
	gorm.Model
	Date       time.Time
	IsChosen   bool
	Note       sql.NullString
	ScheduleID uint
}
