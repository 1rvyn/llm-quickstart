package routes

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/1rvyn/llm-quickstart/frontend/models"
	"github.com/gofiber/fiber/v2"
)

func UploadPage(c *fiber.Ctx) error {

	labels := &models.Labels{}

	file, err := os.Open("labels.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, labels)

	if err != nil {
		log.Fatal(err)
	}

	return c.Render("upload", fiber.Map{
		"Title":  "Upload new PDF?",
		"Labels": labels.Labels,
	})

}
