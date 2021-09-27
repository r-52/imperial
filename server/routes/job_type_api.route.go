package routes

import "github.com/gofiber/fiber/v2"

func getJobTypeById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobTypes(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteJobType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateJobType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createJobType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupJobTypeRoutes(router fiber.Router) {
	router.Get("/:company/:id", getJobTypeById)
	router.Get("/:company/", getAllJobTypes)
	router.Delete("/:company/:id", deleteJobType)
	router.Patch("/:company/update/:id", updateJobType)
	router.Post("/:company/", createJobType)
}
