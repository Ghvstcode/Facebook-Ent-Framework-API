package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/GhvstCode/LR-ENT/handlers"
)

func Routes(app *fiber.App){
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/createnote", handlers.CreateNote)
	api.Get("/readnotes", handlers.ReadNote)
	api.Get("/searchnotes/:title", handlers.SearchNotes)
	api.Put("/updatenote/:id", handlers.UpdateNote)
	api.Delete("/deletenote/:id", handlers.DeleteNotes)
}

func main() {
	app := fiber.New()

	Routes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Unable to start server")
	}

	fmt.Println("Server is up on port 3000")
}