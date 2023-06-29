package routes

import "github.com/gofiber/fiber/v2"

func Ingest(c *fiber.Ctx) error {
	return c.SendString("Ingest")
}
