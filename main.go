package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"os"
	"yusa/seed"
	"yusa/todo"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	app := fiber.New()
	app.Use(cors.New())

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5437", "postgres", "postgres", "postgres"))
	if err != nil {
		return err
	}
	if err = seed.MigrateTables(db); err != nil {
		fmt.Println("error occurred when migrating tables")
		return err
	}

	repository := todo.NewRepository(db)
	service := todo.NewService(repository)
	handler := todo.NewHandler(service)

	app.Get("/todos", handler.GetTodos)
	app.Post("/todos", handler.AddTodo)

	return app.Listen(":7575")
}
