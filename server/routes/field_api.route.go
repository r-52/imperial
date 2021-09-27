package routes

import "github.com/gofiber/fiber/v2"

func getFieldById(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func getAllFields(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func deleteField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func updateField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func createField(c *fiber.Ctx) error {
	return c.SendString("NOT_IMPLEMENTED")
}

func SetupFieldRoutes(router fiber.Router) {
	router.Get("/:id", getFieldById)
	router.Get("/", getAllFields)
	router.Delete("/:id", deleteField)
	router.Patch("/update/:id", updateField)
	router.Post("/", createField)
}
