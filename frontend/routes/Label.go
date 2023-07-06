package routes

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/1rvyn/llm-quickstart/frontend/models"
	"github.com/gofiber/fiber/v2"
)

func Label(c *fiber.Ctx) error {
	// Parse the request body
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	fmt.Println(body)

	// Read the existing labels
	labels := &models.Labels{}
	file, err := os.Open("labels.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	defer file.Close()

	// Unmarshal the JSON data
	err = json.NewDecoder(file).Decode(labels)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Generate a new ID for the label
	newID := strconv.Itoa(len(labels.Labels) + 1)

	fmt.Println("new label id is: ", newID)

	// Add the new label
	labels.Labels = append(labels.Labels, models.Label{
		ID:   newID,
		Name: body["label"],
	})

	// Marshal the updated labels back to JSON
	updatedData, err := json.Marshal(labels)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile("labels.json", updatedData, 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Send a success response
	return c.JSON(fiber.Map{
		"success": true,
	})
}
