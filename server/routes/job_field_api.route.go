package routes

import "github.com/gofiber/fiber/v2"

func getJobFieldById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobFieldsForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobFieldsForJobAndCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteJobField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateJobField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createJobField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupJobFieldRoutes(router fiber.Router) {
	router.Get("/:company/:id", getJobFieldById)
	router.Get("/:company/", getAllJobFieldsForCompany)
	router.Delete("/:company/:id", deleteJobField)
	router.Patch("/updated/:company/:id", updateJobField)
	router.Post("/:company/", createJobField)
	router.Get("/job/:job/:company/", getAllJobFieldsForJobAndCompany)
}
