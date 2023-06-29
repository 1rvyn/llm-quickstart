package routes

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Upload(c *fiber.Ctx) error {
	fmt.Printf("Upload called\n")

	id := c.Params("id")
	fmt.Printf("ID from URL: %s\n", id)

	var result []string

	// Loop over the expected file keys
	for i := 0; ; i++ {
		fileKey := fmt.Sprintf("file%d", i)
		file, err := c.FormFile(fileKey)

		// Break the loop if a file is not found
		if err != nil {
			break
		}

		// Create a unique file name
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
		// Save the file to disk
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileName)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server error",
			})
		}
		result = append(result, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	}

	// If no file was uploaded
	if len(result) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No file was uploaded",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": result,
	})
}

// TODO: Make this add the uplaoded files to the ID folder
// - Make it run ingest.py by passing the ID to the python script (will need to adjust ingest.py to handle param from stdin)
