package routes

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/1rvyn/llm-quickstart/frontend/database"
	"github.com/1rvyn/llm-quickstart/frontend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Search(c *fiber.Ctx) error {
	// parse the search query using the model

	fmt.Println("A new search query just got asked")

	var newQuestion models.Question

	err := c.BodyParser(&newQuestion)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Error parsing search query",
		})
	}

	fmt.Println("A new question just got asked", newQuestion.Question)

	// parse the users cookie from the request

	cookie := c.Cookies("session")

	fmt.Println("The user's cookie is", cookie)

	// extract the claims from the cookie

	// TODO: Set and use a proper secret as a env variable
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // replace "secret" with your own secret key (in login.go it is 'secret')
	})

	if err != nil {
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

	// use the claims as needed

	fmt.Println("The user's ID is", claims["iss"])

	userID, err := strconv.Atoi(claims["iss"].(string))
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	newQuestion.UserID = uint(userID)

	fmt.Println(newQuestion)

	// Use os/exec to run Python script
	cmd := exec.Command("python3", "/Users/irvyn/work/chat-pdf/src/single-pdf.py")
	cmd.Dir = "/Users/irvyn/work/chat-pdf/src" // <-- set the working directory here
	// cmd.Stdin = strings.NewReader(newQuestion.Question)            // Pass the question as an input to the script
	cmd.Stdin = strings.NewReader(newQuestion.Question) // ask question
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // Capture the output of the script
	cmd.Stderr = &stderr // Capture the standard error output of the script
	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Error running Python script",
			"error":   err.Error(),
			"stderr":  stderr.String(),
		})
	}

	// add answer to newQuestion
	newQuestion.Answer = out.String()

	// add the question to the database
	if err := database.Database.Db.Create(&newQuestion).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Search successful",
		"output":  out.String(),
	})
}
