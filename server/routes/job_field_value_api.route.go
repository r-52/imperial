package routes

import "github.com/gofiber/fiber/v2"

func getJobFieldValueById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobFieldValuesForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteJobFieldValue(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateJobFieldValue(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createJobFieldValue(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobFieldValuesForJobAndCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupJobFieldValueRoutes(router fiber.Router) {
	router.Get("/:company/:id", getJobFieldValueById)
	router.Get("/:company/", getAllJobFieldValuesForCompany)
	router.Delete("/:company/:id", deleteJobFieldValue)
	router.Patch("/updated/:company/:id", updateJobFieldValue)
	router.Post("/:company/", createJobFieldValue)
	router.Get("/job/:job/:company/", getAllJobFieldValuesForJobAndCompany)
}
