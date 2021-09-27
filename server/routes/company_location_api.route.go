package routes

import "github.com/gofiber/fiber/v2"

func getCompanyLocationById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllCompanyLocations(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupCompanyLocationRoutes(router fiber.Router) {
	router.Get("/:company/:id", getCompanyLocationById)
	router.Get("/:company/", getAllCompanyLocations)
	router.Delete("/:company/:id", deleteCompanyLocation)
	router.Patch("/:company/update/:id", updateCompanyLocation)
	router.Post("/:company/", createCompanyLocation)
}
