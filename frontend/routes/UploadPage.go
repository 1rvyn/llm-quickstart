package routes

import "github.com/gofiber/fiber/v2"

func UploadPage(c *fiber.Ctx) error {
	return c.Render("upload", fiber.Map{
		"Title": "Upload new PDF?",
	})

}
