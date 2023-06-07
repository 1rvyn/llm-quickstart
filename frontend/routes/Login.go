package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// get the username and password from the request body
	var logindata struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&logindata); err != nil {
		return err
	}

	// debug print
	fmt.Println(logindata)

	// return a success message
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful!",
	})

}
