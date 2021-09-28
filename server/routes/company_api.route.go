package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
	"github.com/r-52/imperial/models"
	companyViewModel "github.com/r-52/imperial/view_model/company"
	errorViewModel "github.com/r-52/imperial/view_model/error"
)

func createCompany(c *fiber.Ctx) error {
	viewModel := new(companyViewModel.CompanyViewModel)
	if err := c.BodyParser(viewModel); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "body parsing failed"}
		return c.JSON(response)
	}

	if _, err := viewModel.Validate(); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "validation failed"}
		return c.JSON(response)
	}

	db := database.GetDatabase()

	company := new(models.Company)
	viewModel.ConvertToDbModel(company)
	db.Create(&company)

	companyPerson := new(models.CompanyPerson)
	companyPerson.Email = viewModel.AdminEmail
	companyPerson.SetRandomPassword()
	companyPerson.SetPasswordResetToken()
	companyPerson.FirstName = "CompanyAdmin"
	companyPerson.LastName = "System"

	companyPerson.CompanyID = company.ID
	db.Create(&companyPerson)

	location := new(models.CompanyLocation)
	location.Zip = viewModel.Zip
	location.City = viewModel.City
	location.Street1 = viewModel.Street1
	location.Street2 = viewModel.Street2
	location.Country = viewModel.Country
	location.County = viewModel.County
	location.Name = viewModel.Name
	location.CompanyID = company.ID
	location.IsMainLocation = true

	db.Create(&location)

	return c.JSON(companyViewModel.ToViewModel(company, location))
}

func updateCompany(c *fiber.Ctx) error {
	viewModel := new(companyViewModel.CompanyViewModel)
	if err := c.BodyParser(viewModel); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "body parsing failed"}
		return c.JSON(response)
	}

	if _, err := viewModel.Validate(); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "validation failed"}
		return c.JSON(response)
	}
	id := c.Params("id")

	company := new(models.Company)
	db := database.GetDatabase()
	db.Where("Uid = ?", id).First(&company)
	viewModel.ConvertToDbModel(company)
	db.Save(&company)
	return c.JSON(company)
}

func deleteCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	company := new(models.Company)
	db := database.GetDatabase()
	db.Where("Uid = ?", id).First(&company)

	if company == nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "delete not possible"}
		return c.JSON(response)
	}

	db.Delete(&company)
	return c.SendStatus(fiber.StatusOK)
}

func getCompanyById(c *fiber.Ctx) error {
	id := c.Params("id")
	company := new(models.Company)
	db := database.GetDatabase()
	err := db.Where("Uid = ?", id).Preload("CompanyLocation").First(&company)
	if err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "company not found"}
		return c.JSON(response)
	}

	for _, v := range company.CompanyLocation {
		if v.IsMainLocation {
			return c.JSON(companyViewModel.ToViewModel(company, &v))
		}
	}
	response := errorViewModel.ErrorViewModel{IsError: true, Message: "no main location found"}
	return c.JSON(response)
}

func SetupCompanyRoutes(router fiber.Router) {
	router.Get("/:id", getCompanyById)
	router.Patch("/update/:id", updateCompany)
	router.Delete("/delete/:id", deleteCompany)
	router.Post("/", createCompany)
}
