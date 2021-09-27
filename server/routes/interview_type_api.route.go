package routes

import "github.com/gofiber/fiber/v2"

func createInterviewType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getInterviewTypeById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func deleteInterviewType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func getAllInterviewTypesForCompany(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func updateInterviewType(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMETED")
}

func SetupInterviewTypeRoutes(router fiber.Router) {
	router.Get("/:company/:id", getInterviewTypeById)
	router.Post("/:company/", createInterviewType)
	router.Delete("/:company/:id", deleteInterviewType)
	router.Get("/:company/", getAllInterviewTypesForCompany)
	router.Patch("/:company/:id", updateInterviewType)
}
