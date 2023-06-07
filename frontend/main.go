package main

import (
	"github.com/1rvyn/llm-quickstart/frontend/database"
	"github.com/1rvyn/llm-quickstart/frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// fiber app

	database.ConnectDb()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./views/public")

	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")

		c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

		return c.Next()
	})

	setupRoutes(app)

	// static files
	app.Static("/", "./public")

	// start server
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Home)
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)
}
