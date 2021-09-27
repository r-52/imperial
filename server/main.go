package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
	"github.com/r-52/imperial/routes"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("/v1")

	company := v1.Group("/company")
	routes.SetupCompanyRoutes(company)

	jobApi := v1.Group("/job")
	routes.SetupJobRoutes(jobApi)

	personApi := v1.Group("/person")
	routes.SetupPersonRoutes(personApi)

	applicationApi := v1.Group("/application")
	routes.SetupApplicationRoutes(applicationApi)

	interviewTypeApi := v1.Group("/interview_type")
	routes.SetupInterviewTypeRoutes(interviewTypeApi)

	companyLocationApi := v1.Group("/company_location")
	routes.SetupCompanyLocationRoutes(companyLocationApi)

	companyPersonApi := v1.Group("/company_person")
	routes.SetupPersonRoutes(companyPersonApi)

	companyPersonJobApi := v1.Group("/company_person_job")
	routes.SetupCompanyPersonJobRoutes(companyPersonJobApi)

	fieldApi := v1.Group("/field")
	routes.SetupFieldRoutes(fieldApi)

	jobCompanyLocation := v1.Group("/job_company_location")
	routes.SetupJobCompanyLocationRoutes(jobCompanyLocation)

	jobFieldApi := v1.Group("/job_field")
	routes.SetupJobFieldRoutes(jobFieldApi)

	jobFieldValueApi := v1.Group("/job_field_value")
	routes.SetupJobFieldValueRoutes(jobFieldValueApi)

	jobTypeApi := v1.Group("/job_type")
	routes.SetupJobTypeRoutes(jobTypeApi)
}

func main() {
	app := fiber.New()

	db := database.InitDatabase()
	db.Exec("SELECT 1;")

	setupRoutes(app)

	app.Static("/", "../web-client/dist/web-client/")

	app.Listen(":3000")
}
