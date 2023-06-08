package main

import (
	"github.com/1rvyn/llm-quickstart/frontend/database"
	"github.com/1rvyn/llm-quickstart/frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	database.ConnectDb() // connect to the database
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views/public")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")

		c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

		return c.Next()
	})

	setupRoutes(app)

	// start server
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Home)
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)
}
