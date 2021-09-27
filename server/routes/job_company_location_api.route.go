package routes

import "github.com/gofiber/fiber/v2"

func getJobCompanyLocationById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllJobCompanyLocations(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteJobCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateJobCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createJobCompanyLocation(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllLocationsForJob(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupJobCompanyLocationRoutes(router fiber.Router) {
	router.Get("/:id", getJobCompanyLocationById)
	router.Get("/", getAllJobCompanyLocations)
	router.Delete("/:id", deleteJobCompanyLocation)
	router.Patch("/update/:id", updateJobCompanyLocation)
	router.Post("/", createJobCompanyLocation)
	router.Get("/job/:job/:company", getAllLocationsForJob)
}
