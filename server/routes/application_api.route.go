package routes

import "github.com/gofiber/fiber/v2"

func createApplication(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getApplicationById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func deleteApplication(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getAllApplicationsForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func updateApplication(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func SetupApplicationRoutes(router fiber.Router) {
	router.Get("/:company/:id", getApplicationById)
	router.Post("/:company/", createApplication)
	router.Delete("/:company/:id", deleteApplication)
	router.Get("/:company/", getAllApplicationsForCompany)
	router.Patch("/:company/:id", updateApplication)
}
