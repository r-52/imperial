package routes

import "github.com/gofiber/fiber/v2"

func getCompanyPersonJobById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllCompanyPersonJobs(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteCompanyPersonJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateCompanyPersonJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createCompanyPersonJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupCompanyPersonJobRoutes(router fiber.Router) {
	router.Get("/:company/:id", getCompanyPersonJobById)
	router.Get("/:company", getAllCompanyPersonJobs)
	router.Delete("/:company/:id", deleteCompanyPersonJob)
	router.Patch("/:company/update/:id", updateCompanyPersonJob)
	router.Post("/:company", createCompanyPersonJob)
}
