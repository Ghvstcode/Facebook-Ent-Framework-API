package handlers

import (
	"context"
	"fmt"
	"log"
	_ "log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/GhvstCode/LR-ENT/database"
	"github.com/GhvstCode/LR-ENT/ent/notes"
)



func CreateNote(c *fiber.Ctx) error{
	note := new(struct{
		Title string
		Content string
		Private bool
	})

	if err := c.BodyParser(&note); err != nil {
		c.Status(400).JSON("Error  Parsing Input")
		return err
	}

	createdNote, err := database.EntClient.Notes.
		Create().
		SetTitle(note.Title).
		SetContent(note.Content).
		SetPrivate(note.Private).
		Save(context.Background())

	if err != nil {
		c.Status(500).JSON("Unable to save note")
		return err
	}

	c.Status(200).JSON(createdNote)

	return nil
}

func ReadNote(c *fiber.Ctx) error{
	readNotes, err := database.EntClient.Notes.
		Query().
		All(context.Background())
	if err != nil {
		c.Status(500).JSON("No Notes Found")
		log.Fatal(err)
	}
	c.Status(200).JSON(readNotes)
	return nil
}

func SearchNotes(c *fiber.Ctx) error{
	query := c.Params("title")
	if query == "" {
		c.Status(400).JSON("No Search Query")
	}
	fmt.Println(query)
	createdNotes, err := database.EntClient.Notes.
		Query().
		Where(notes.Title(query)).
		All(context.Background())

	if err != nil {
		c.Status(500).JSON("No Notes Found")
		log.Fatal(err)
	}

	c.Status(200).JSON(createdNotes)
	return nil
}

func UpdateNote(c *fiber.Ctx) error{
	note := new(struct{
		Content string
	})

	if err := c.BodyParser(&note); err != nil {
		c.Status(400).JSON("Error  Parsing Input")
		return err
	}

	idParam := c.Params("id")
	if idParam == "" {
		c.Status(400).JSON("No Search Query")
	}

	id, _ := strconv.Atoi(idParam)
	UpdatedNotes, err := database.EntClient.Notes.
		UpdateOneID(id).
		SetContent(note.Content).
		Save(context.Background())

	if err != nil {
		c.Status(500).JSON("No Notes Found")
		log.Fatal(err)
	}

	c.Status(200).JSON(UpdatedNotes)
	return nil
}

func DeleteNotes(c *fiber.Ctx) error{
	idParam := c.Params("id")
	if idParam == "" {
		c.Status(400).JSON("No Search Query")
	}

	id, _ := strconv.Atoi(idParam)

	err := database.EntClient.Notes.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		c.Status(500).JSON("Unable to Perform Operation")
	}

	c.Status(200).JSON("Success")

	return nil
}

//curl --location --request POST 'http://localhost:3000/api/v1/createnote' \
//--header 'Content-Type: application/json' \
//--data-raw '{
//"title": "A poem about Lorem",
//"content": "Lorem Ipsum something something"
//}'

//curl --location --request GET 'http://localhost:3000/api/v1/readnote/:important' \
//--header 'Content-Type: application/json'
