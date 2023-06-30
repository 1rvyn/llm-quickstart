package routes

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func DeleteFiles(c *fiber.Ctx) error {
	type RequestBody struct {
		Files    []string `json:"files"`
		FolderId string   `json:"folderId"`
	}

	var body RequestBody

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	for _, file := range body.Files {
		// Construct the file path
		filePath := fmt.Sprintf("/Users/irvyn/work/chat-pdf/src/src/data/chroma/%s/%s", body.FolderId, file)

		err := os.Remove(filePath)
		if err != nil {
			fmt.Printf("Error deleting file at path: %s. Error: %s\n", filePath, err)
			return c.Status(500).SendString(err.Error())
		}
	}

	return c.JSON(fiber.Map{
		"message": "Files deleted successfully",
	})
}
