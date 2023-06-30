package routes

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FolderContents(c *fiber.Ctx) error {
	id := c.Params("id")

	// Construct the folder path
	folderPath := fmt.Sprintf("/Users/irvyn/work/chat-pdf/src/src/data/chroma/%s/", id)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Filter PDF files
	var pdfFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".pdf") {
			pdfFiles = append(pdfFiles, file.Name())
		}
	}

	return c.JSON(pdfFiles)
}
