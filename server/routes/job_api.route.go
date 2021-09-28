package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
	"github.com/r-52/imperial/models"
	errorViewModel "github.com/r-52/imperial/view_model/error"
	jobViewModel "github.com/r-52/imperial/view_model/job"
)

func getJobById(c *fiber.Ctx) error {
	companyUid := c.Get("company", "")
	jobUid := c.Get("id", "")
	if len(jobUid) == 0 || len(companyUid) == 0 {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "company or job not provided"}
		return c.JSON(response)
	}

	db := database.GetDatabase()
	company := new(models.Company)
	job := new(models.Job)

	if err := db.Find("uid = ?", companyUid).First(&company); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "company not found"}
		return c.JSON(response)
	}
	if err := db.Find("uid = ? AND company_id = ?", jobUid, company.ID).First(&job); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "job not found"}
		return c.JSON(response)
	}

	return c.JSON(job)
}

func getAllJobsForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createJob(c *fiber.Ctx) error {
	viewModel := new(jobViewModel.JobViewModel)
	if err := c.BodyParser(viewModel); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "body parsing failed"}
		return c.JSON(response)
	}

	if _, err := viewModel.Validate(); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "validation failed"}
		return c.JSON(response)
	}

	db := database.GetDatabase()

	companyUid := c.Get("company", "")
	company := new(models.Company)
	if err := db.Find("uid = ?", companyUid).First(&company); err != nil {
		response := errorViewModel.ErrorViewModel{IsError: true, Message: "company was not found"}
		return c.JSON(response)
	}

	job := viewModel.ConvertToDbModel()
	company.Job = append(company.Job, job)

	db.Save(&company)

	// create the person mapping
	for _, v := range viewModel.Persons {
		person := new(models.CompanyPerson)
		if err := db.Where("uid = ?", v).First(&person); err != nil {
			log.Println(fmt.Println("got the person id but no person was found in the db %i", v))
			continue
		}

		companyPerson := new(models.CompanyPersonJob)
		companyPerson.CompanyPersonID = person.ID
		companyPerson.JobID = job.ID
		job.CompanyPersonJob = append(job.CompanyPersonJob, companyPerson)
	}

	// create the location mapping
	for _, v := range viewModel.CompanyLocations {
		location := new(models.CompanyLocation)
		if err := db.Where("uid = ?", v).First(&location); err != nil {
			log.Println(fmt.Println("got the location id but no location was found in the db %i", v))
			continue
		}
		companyLocation := new(models.JobCompanyLocation)
		companyLocation.CompanyLocationID = location.ID
		companyLocation.JobID = job.ID
		job.JobCompanyLocation = append(job.JobCompanyLocation, companyLocation)
	}

	db.Save(&job)

	viewModel.Uid = job.Uid
	return c.JSON(viewModel)
}

func SetupJobRoutes(router fiber.Router) {
	router.Get("/:company/:id", getJobById)
	router.Get("/:company/", getAllJobsForCompany)
	router.Delete("/:company/:id", deleteJob)
	router.Patch("/update/:company/:id", updateJob)
	router.Post("/:company/", createJob)
}
