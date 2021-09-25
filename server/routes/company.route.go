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
	viewModel := new(company.Company)
	if err := c.BodyParser(viewModel); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "body parsing failed"}
		return c.JSON(response)
	}

	if _, err := viewModel.Validate(); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "validation failed"}
		return c.JSON(response)
	}

	company := new(models.Company)
	company.AdminEmail = viewModel.AdminEmail
	company.Name = viewModel.Name
	company.Imprint = sql.NullString{String: viewModel.Imprint, Valid: true}
	company.CompanyUrl = viewModel.Url
	company.CustomLogoUrl = sql.NullString{String: viewModel.CustomLogoUrl, Valid: true}
	company.DataProtectionTemplate = sql.NullString{String: viewModel.DataProtectionTemplate, Valid: true}
	company.TermsOfServiceTemplate = sql.NullString{String: viewModel.TermsOfServiceTemplate, Valid: true}
	company.CustomStyleSheetInjection = sql.NullString{String: viewModel.CustomStyleSheetInjection, Valid: true}

	db := database.GetDatabase()
	db.Create(&company)

	return c.SendString("NOT_IMPLEMENTED")
}

func updateCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getCompanyById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllCompanies(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupCompanyRoutes(router fiber.Router) {
	router.Get("/:id", getCompanyById)
	router.Get("/", getAllCompanies)
	router.Patch("/update/:id", updateCompany)
	router.Delete("/delete/:id", deleteCompany)
	router.Post("/", createCompany)
}
