package routes

import "github.com/gofiber/fiber/v2"

func getJobById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
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
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupJobRoutes(router fiber.Router) {
	router.Get("/:company/:id", getJobById)
	router.Get("/:company/", getAllJobsForCompany)
	router.Delete("/:company/:id", deleteJob)
	router.Patch("/updated/:company/:id", updateJob)
	router.Post("/:company/", createJob)
}
