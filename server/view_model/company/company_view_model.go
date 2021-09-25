package company

import "github.com/go-playground/validator/v10"

type Company struct {
	Name    string `json:"name" validate:"required,min=2,max=128"`
	Imprint string `json:"imprint" validate:"required,min=2,max=1024"`

	AdminEmail string `json:"adminEmail" validate:"required,email"`
	Url        string `json:"url" validate:"required,url"`

	Street1 string `json:"street1" validate:"required,min=2,max=1024"`
	Street2 string `json:"street2"`
	Zip     string `json:"zip" validate:"required,min=2,max=1024"`
	Country string `json:"country" validate:"required,min=2,max=1024"`
	County  string `json:"county" validate:"required,min=2,max=1024"`

	CustomStyleSheetInjection string `json:"customStyleSheet"`
	CustomLogoUrl             string `json:"customLogoUrl"`

	DataProtectionTemplate string `json:"dataProtectionTemplate"`
	TermsOfServiceTemplate string `json:"TermsOfServiceTemplate"`
}

func (c *Company) Validate() (bool, error) {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}
