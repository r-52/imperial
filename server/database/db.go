package database

import (
	"github.com/r-52/imperial/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	return InitDatabase()
}

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("could not open database!")
	}

	db.AutoMigrate(&models.Company{},
		&models.CompanyLocation{},
		&models.CompanyPerson{},
		&models.Knowledge{},
		&models.JobType{},
		&models.CompanyPersonJob{},
		&models.Field{},
		&models.Job{},
		&models.JobCompanyLocation{},
		&models.JobField{},
		&models.JobFieldValue{},
		&models.ApplicationStatus{},
		&models.Application{},
		&models.Schedule{},
		&models.ScheduleEntry{},
		&models.MailTemplate{},
		&models.InterviewType{},
		&models.ApplicationInterviewNote{})

	return db
}
