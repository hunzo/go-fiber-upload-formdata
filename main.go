package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "test",
		})
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		username := c.FormValue("username")
		if username == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error FormValue username",
			})
		}
		passphrase := c.FormValue("passphrase")
		if passphrase == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "error FormValue passphrase",
			})
		}

		destination := fmt.Sprintf("./uploads/%s"+".temp", file.Filename)
		if err := c.SaveFile(file, destination); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"fileName":   file.Filename,
			"username":   username,
			"passphrase": passphrase,
		})
	})

	log.Println(app.Listen(":8080"))
}
