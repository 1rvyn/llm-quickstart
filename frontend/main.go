package main

import (
	"fmt"
	"strconv"

	"github.com/1rvyn/llm-quickstart/frontend/database"
	"github.com/1rvyn/llm-quickstart/frontend/models"
	"github.com/1rvyn/llm-quickstart/frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/golang-jwt/jwt"
)

func main() {

	database.ConnectDb() // connect to the database
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 100 * 1024 * 1024, // 100 MB

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
	app.Use(checkCookieStatus())

	app.Get("/", routes.Home)
	app.Get("/login", routes.LoginPage)
	app.Post("/api/login", routes.Login)
	app.Post("/api/register", routes.Register) // temp
	app.Get("/register", routes.RegisterPage)

	protected := app.Group("/")
	protected.Use(requireAuth())
	protected.Get("/search", routes.SearchPage)
	protected.Post("/api/search", routes.Search)

	admin := app.Group("/admin")
	admin.Use(requireAdminAuth())
	admin.Post("/api/ingest/:id", routes.Ingest)
	admin.Get("/upload", routes.UploadPage)
	admin.Post("/api/upload/:id", routes.Upload)
}

func checkCookieStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loggedIn := false

		cookie := c.Cookies("session")
		fmt.Println(cookie)
		if cookie != "" {
			// Check if the session cookie is valid

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("salty"), nil // TODO replace all uses of these keys with proper env vars
			})

			if err != nil {
				fmt.Println("error", err)
				return c.JSON(fiber.Map{
					"success": false,
					"message": "Error extracting claims from cookie",
				})
			}

			if !token.Valid {
				return c.JSON(fiber.Map{
					"success": false,
					"message": "Invalid token",
				})
			}

			// If the cookie is valid, set loggedIn to true
			loggedIn = true
		}

		c.Locals("LoggedIn", loggedIn)

		return c.Next()
	}
}

func requireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("middleware ran: ")
		cookie := c.Cookies("session")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("salty"), nil // TODO replace all uses of these keys with proper env vars
		})

		// TODO: Change this to a redirect to home page with an error message that gets displayed
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Error extracting claims from cookie",
			})
		}

		if !token.Valid {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Invalid token",
			})
		}

		loggedIn := true
		c.Locals("LoggedIn", loggedIn)

		userID, err := strconv.Atoi(claims["iss"].(string))
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Error parsing user ID",
				"error":   err.Error(),
			})
		}

		fmt.Println("User ID is", userID)

		var userAsking models.Accounts

		// get the user

		if err := database.Database.Db.Find(&userAsking, "id = ?", uint(userID)).Error; err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"Error":   "Error finding user",
			})
		}

		return c.Next()
	}
}

func requireAdminAuth() fiber.Handler {
	fmt.Println("admin middleware ran: ")

	// TODO: Same as above but we will check the user role to see if they are an admin

	return func(c *fiber.Ctx) error {
		fmt.Println("middleware ran: ")
		cookie := c.Cookies("session")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("salty"), nil // TODO replace all uses of these keys with proper env vars
		})

		// TODO: Change this to a redirect to home page with an error message that gets displayed
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Error extracting claims from cookie",
			})
		}

		if !token.Valid {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Invalid token",
			})
		}

		userID, err := strconv.Atoi(claims["iss"].(string))
		if err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Error parsing user ID",
				"error":   err.Error(),
			})
		}

		fmt.Println("User ID is", userID)

		var userAsking models.Accounts

		// get the user

		if err := database.Database.Db.Find(&userAsking, "id = ?", uint(userID)).Error; err != nil {
			fmt.Println(err)
			return c.JSON(fiber.Map{
				"success": false,
				"Error":   "Error finding user",
			})
		}

		if userAsking.UserRole != 15 {
			return c.JSON(fiber.Map{
				"success": false,
				"Error":   "Not authorized",
			})
		}

		fmt.Println("User is an admin")

		return c.Next()
	}
}
