package routes

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	loggedIn := false

	cookie := c.Cookies("session")
	if cookie != "" {
		// Check if the session cookie is valid
		// ... validation logic ...
		// If the cookie is valid, set loggedIn to true
		loggedIn = true
	}

	return c.Render("home", fiber.Map{
		"Title":    "time to help",
		"LoggedIn": loggedIn,
	})
}
