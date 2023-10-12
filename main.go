package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"todo/ent"
	"todo/module/todo"
	"todo/module/user"
)

func main() {
	app := fiber.New(*getConfig())
	app.Use(cors.New(*getCorsConfig()))

	db := dbConnect()

	user.Routes(app.Group("/user"), db)
	todo.Routes(app.Group("/"), db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

func getConfig() *fiber.Config {
	return &fiber.Config{
		Prefork: false,
		AppName: "ToDo",
	}
}

func getCorsConfig() *cors.Config {
	return &cors.Config{
		AllowCredentials: true,
	}
}

func dbConnect() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo password=1234 sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	return client
}
