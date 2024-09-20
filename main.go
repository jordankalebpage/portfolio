package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
	Deleted   bool   `json:"deleted"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Get("/api/todos", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(todos)
	})

	app.Post("/api/todos", func(c fiber.Ctx) error {
		todo := Todo{}
		if err := c.Bind().Body(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		todo := Todo{}

		if err := c.Bind().Body(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
		}

		for i, t := range todos {
			if fmt.Sprint(t.ID) == id {
				todos[i] = todo
				return c.Status(fiber.StatusOK).JSON(todos[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "todo not found"})
	})

	app.Delete("/api/todos/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for i, t := range todos {
			if fmt.Sprint(t.ID) == id {
				todo := todos[i]
				todo.Deleted = true
				todos[i] = todo

				return c.Status(fiber.StatusOK).JSON(todo)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "todo not found"})
	})

	log.Fatal(app.Listen(":8080"))
}
