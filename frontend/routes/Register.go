package routes

import (
	"fmt"

	"github.com/1rvyn/llm-quickstart/frontend/database"
	"github.com/1rvyn/llm-quickstart/frontend/models"
	"github.com/1rvyn/llm-quickstart/frontend/utils"
	"github.com/gofiber/fiber/v2"
)

var SALT = "salty"

func Register(c *fiber.Ctx) error {
	var registerData map[string]string

	if err := c.BodyParser(&registerData); err != nil {
		return err
	}

	hashedPassword := make(chan []byte) // channel to recieve the hashed password

	go func() {
		hashedPassword <- utils.HashPassword(registerData["password"], []byte(SALT))
		close(hashedPassword)
	}()

	encryptedPassword := <-hashedPassword

	if encryptedPassword == nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "could not encrypt password",
		})
	}

	// TODO: hash the password + send email verification

	user := models.Users{
		Username: registerData["username"],
		Password: encryptedPassword,
	}

	if err := database.Database.Db.Create(&user).Error; err != nil {
		return err
	}

	fmt.Println("Registered this: ", registerData)

	return c.JSON(fiber.Map{
		"message": "successfully registered",
		"User":    user,
	})
}
