package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
	"github.com/r-52/imperial/models"
	"github.com/r-52/imperial/view_model/company"
	errorViewModel "github.com/r-52/imperial/view_model/error"
)

func createCompany(c *fiber.Ctx) error {
	viewModel := new(company.CompanyViewModel)
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
	location.Street2 = sql.NullString{String: viewModel.Street2, Valid: true}
	location.Country = viewModel.Country
	location.County = sql.NullString{String: viewModel.County, Valid: true}
	location.Name = viewModel.Name
	location.CompanyID = company.ID
	location.IsMainLocation = true

	db.Create(&location)

	return c.JSON(company)
}

func updateCompany(c *fiber.Ctx) error {
	viewModel := new(company.CompanyViewModel)
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
	db.Where("Uid = ?", id).First(&company)
	if company == nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "delete not possible"}
		return c.JSON(response)
	}
	return c.JSON(company)
}

func SetupCompanyRoutes(router fiber.Router) {
	router.Get("/:id", getCompanyById)
	router.Patch("/update/:id", updateCompany)
	router.Delete("/delete/:id", deleteCompany)
	router.Post("/", createCompany)
}
