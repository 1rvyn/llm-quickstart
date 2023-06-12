package routes

import "github.com/gofiber/fiber/v2"

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}
