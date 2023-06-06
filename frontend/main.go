package main

import "github.com/gofiber/fiber/v2"

func main() {
	// fiber app
	app := fiber.New()

	setupRoutes(app)

	// static files
	app.Static("/", "./public")

	// start server
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Home)
}
