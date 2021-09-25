package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
)

func main() {
	app := fiber.New()

	db := database.InitDatabase()
	db.Exec("SELECT 1;")

	app.Static("/", "../web-client/dist/web-client/")

	app.Listen(":3000")
}
