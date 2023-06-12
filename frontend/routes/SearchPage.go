package routes

import "github.com/gofiber/fiber/v2"

func SearchPage(c *fiber.Ctx) error {
	return c.Render("search", fiber.Map{
		"Title": "Do you need help using our Simulator?",
	})
}
