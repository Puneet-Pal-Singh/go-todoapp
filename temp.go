package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Todo1 struct {
    ID int `json:"id"`
    Completed bool `json:"completed"`
    Body string `json:"body"`
}

func main1() {
    // Initialize a new Fiber app
    app := fiber.New()
	fmt.Println("Hello words")

    // dotenv
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file: %v", err)
    }

    // Retrieve the environment variables
    // databaseURL := os.Getenv("DATABASE_URL")
    // apiKey := os.Getenv("API_KEY")
    PORT := os.Getenv("PORT")

    todos := []Todo{}

    // Define a route for the GET method on the root path '/'
    app.Get("/api/todos", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.Status(200).JSON(todos)
        // return c.Status(200).JSON(fiber.Map{
        //     "msg":"Hello Worlddd",
        // })
        // return c.SendString("Hello, World 👋!")
    })

    app.Post("/api/todos", func(c fiber.Ctx) error {
        todo := new(Todo)

        if err := c.Bind().Body(todo); err != nil {
            return err
        }

        if todo.Body == ""{
            return c.Status(400).JSON(fiber.Map{
                "error": "Todo body is required",
            })
        }

        // todo.ID = len(todos) + 1
        todos = append(todos, *todo)

        return c.Status(201).JSON(todo)
    })

    // Update a Todo
    app.Patch("/api/todos/:id", func (c fiber.Ctx) error {
        id := c.Params("id")

        for i, todo := range todos {
            if fmt.Sprint(todo.ID) == id {
                // updating completed field
                todos[i].Completed = true
                return c.Status(200).JSON(todos[i])
            }
        }
        return c.Status(404).JSON(fiber.Map{
            "error": "Todo not Found", 
        })
    })

    // Delete a Todo

    app.Delete("api/todos/:id", func(c fiber.Ctx) error {
        
        id := c.Params("id")

        for i, todo := range todos{
            if fmt.Sprint(todo.ID) == id {
                todos = append(todos[:i], todos[i+1:]...)
                return c.Status(200).JSON(fiber.Map{
                    "success": "true",
                })
            }
        }
        return c.Status(404).JSON(fiber.Map{"error": "TODO not found"})
    })

    // Start the server on port 4000
    log.Fatal(app.Listen(":" + PORT))
}