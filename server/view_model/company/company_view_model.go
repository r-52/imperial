package company

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/r-52/imperial/models"
)

type CompanyViewModel struct {
	Uid     string `json:"uid"`
	Name    string `json:"name" validate:"required,min=2,max=128"`
	Imprint string `json:"imprint" validate:"required,min=2,max=1024"`

	AdminEmail string `json:"adminEmail" validate:"required,email"`
	Url        string `json:"url" validate:"required,url"`

	Street1 string `json:"street1" validate:"required,min=2,max=1024"`
	Street2 string `json:"street2"`
	Zip     string `json:"zip" validate:"required,min=2,max=1024"`
	City    string `json:"city" validate:"required,min=2,max=1024"`
	Country string `json:"country" validate:"required,min=2,max=1024"`
	County  string `json:"county"`

	CustomStyleSheetInjection string `json:"customStyleSheet"`
	CustomLogoUrl             string `json:"customLogoUrl"`

	DataProtectionTemplate string `json:"dataProtectionTemplate"`
	TermsOfServiceTemplate string `json:"TermsOfServiceTemplate"`
}

func (c *CompanyViewModel) Validate() (bool, error) {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CompanyViewModel) ConvertToDbModel(company *models.Company) {
	company.AdminEmail = c.AdminEmail
	company.Name = c.Name
	company.Imprint = sql.NullString{String: c.Imprint, Valid: true}
	company.CompanyUrl = c.Url
	company.CustomLogoUrl = sql.NullString{String: c.CustomLogoUrl, Valid: true}
	company.DataProtectionTemplate = sql.NullString{String: c.DataProtectionTemplate, Valid: true}
	company.TermsOfServiceTemplate = sql.NullString{String: c.TermsOfServiceTemplate, Valid: true}
	company.CustomStyleSheetInjection = sql.NullString{String: c.CustomStyleSheetInjection, Valid: true}
}
