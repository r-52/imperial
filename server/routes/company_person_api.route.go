package routes

import "github.com/gofiber/fiber/v2"

func createPerson(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getPersonById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func deletePerson(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getAllPersonsForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func updatePerson(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func SetupPersonRoutes(router fiber.Router) {
	router.Get("/:company/:id", getPersonById)
	router.Post("/:company/", createPerson)
	router.Delete("/:company/:id", deletePerson)
	router.Get("/:company/", getAllPersonsForCompany)
	router.Patch("/:company/:id", updatePerson)
}
