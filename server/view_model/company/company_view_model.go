package company

import (
	"github.com/go-playground/validator/v10"
	"github.com/r-52/imperial/models"
)

type CompanyViewModel struct {
	Uid     string `json:"uid,omitempty"`
	Name    string `json:"name,omitempty" validate:"required,min=2,max=128"`
	Imprint string `json:"imprint,omitempty" validate:"required,min=2,max=1024"`

	AdminEmail string `json:"adminEmail,omitempty" validate:"required,email"`
	Url        string `json:"url,omitempty" validate:"required,url"`

	Street1 string `json:"street1" validate:"required,min=2,max=1024"`
	Street2 string `json:"street2,omitempty"`
	Zip     string `json:"zip,omitempty" validate:"required,min=2,max=1024"`
	City    string `json:"city,omitempty" validate:"required,min=2,max=1024"`
	Country string `json:"country,omitempty" validate:"required,min=2,max=1024"`
	County  string `json:"county,omitempty"`

	CustomStyleSheetInjection string `json:"customStyleSheet"`
	CustomLogoUrl             string `json:"customLogoUrl"`

	DataProtectionTemplate string `json:"dataProtectionTemplate,omitempty"`
	TermsOfServiceTemplate string `json:"TermsOfServiceTemplate,omitempty"`
}

func ToViewModel(c *models.Company, l *models.CompanyLocation) CompanyViewModel {
	viewModel := new(CompanyViewModel)
	viewModel.AdminEmail = c.AdminEmail
	viewModel.Uid = c.Uid
	viewModel.Name = c.Name
	viewModel.Imprint = c.Imprint
	viewModel.Url = c.CompanyUrl
	viewModel.CustomLogoUrl = c.CustomLogoUrl
	viewModel.CustomStyleSheetInjection = c.CustomStyleSheetInjection
	viewModel.DataProtectionTemplate = c.DataProtectionTemplate
	viewModel.TermsOfServiceTemplate = c.TermsOfServiceTemplate

	viewModel.Street1 = l.Street1
	viewModel.Street2 = l.Street2
	viewModel.Country = l.Country
	viewModel.County = l.County
	viewModel.City = l.City
	viewModel.Zip = l.Zip

	return *viewModel
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
	company.Imprint = c.Imprint
	company.CompanyUrl = c.Url
	company.CustomLogoUrl = c.CustomLogoUrl
	company.DataProtectionTemplate = c.DataProtectionTemplate
	company.TermsOfServiceTemplate = c.TermsOfServiceTemplate
	company.CustomStyleSheetInjection = c.CustomStyleSheetInjection
}
