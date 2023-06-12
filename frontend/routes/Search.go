package routes

import "github.com/gofiber/fiber/v2"

func Search(c *fiber.Ctx) error {
	// TODO: we validate the search -
	// we store who searched what and then perform the search and save the result

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Search successful",
	})
}
