package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/r-52/imperial/database"
	"github.com/r-52/imperial/routes"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("/v1")

	company := v1.Group("/company")
	routes.SetupCompanyRoutes(company)
}

func main() {
	app := fiber.New()

	db := database.InitDatabase()
	db.Exec("SELECT 1;")

	setupRoutes(app)

	app.Static("/", "../web-client/dist/web-client/")

	app.Listen(":3000")
}
