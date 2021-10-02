package models

import (
	"database/sql"
	"errors"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyPerson struct {
	gorm.Model

	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	AvatarUrl    sql.NullString `json:"avatarUrl"`
	PasswordHash string         `json:"-"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Position     string         `json:"position"`

	PasswordResetToken      string       `json:"-"`
	PasswordResetTokenSetAt sql.NullTime `json:"-"`

	Uid string `json:"uid" gorm:"uniqueIndex"`

	CompanyID                uint                        `json:"company"`
	CompanyPersonJob         []*CompanyPersonJob         `json:"companyPersonJob"`
	AcceptedApplication      []*Application              `json:"acceptedApplications" gorm:"foreignKey:AcceptedByPersonID"`
	RejectedApplication      []*Application              `json:"rejectedApplications" gorm:"foreignKey:RejectedByPersonID"`
	InvitedApplication       []*Application              `json:"invitedApplications" gorm:"foreignKey:InvitedByPersonID"`
	Schedule                 []*Schedule                 `json:"schedule"`
	ApplicationInterviewNote []*ApplicationInterviewNote `json:"applicationInterviewNote"`
}

func (c *CompanyPerson) SetPasswordResetToken() {
	id, err := gonanoid.New(64)
	if err != nil {
		panic("could not generate random token")
	}
	c.PasswordResetToken = id
	c.PasswordResetTokenSetAt = sql.NullTime{Time: time.Now(), Valid: true}
}

func (c *CompanyPerson) SetRandomPassword() {
	id, err := gonanoid.New(64)
	if err != nil {
		panic("could not generate random password")
	}
	hash, err := c.HashPassword(id)
	if err != nil {
		panic("could not hash password")
	}
	c.PasswordHash = hash
}

func (c *CompanyPerson) HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

func (c *CompanyPerson) IsPasswordMatch(plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(c.PasswordHash), []byte(plainPassword))
	return err == nil, err
}

func (c *CompanyPerson) BeforeCreate(tx *gorm.DB) (err error) {
	id, e := gonanoid.New(32)

	if e != nil {
		err = errors.New("could not generate unique id")
	}

	c.Uid = id
	return
}
