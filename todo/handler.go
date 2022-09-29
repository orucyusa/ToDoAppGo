package todo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type TodoService interface {
	GetTodos() ([]Todo, error)
	AddTodo(content string) error
}
type FiberHandler struct {
	service TodoService
}

func NewHandler(service TodoService) *FiberHandler {
	return &FiberHandler{service: service}
}
func (f *FiberHandler) GetTodos(ctx *fiber.Ctx) error {

	todos, err := f.service.GetTodos()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(todos)
}
func (f *FiberHandler) AddTodo(ctx *fiber.Ctx) error {
	content := ctx.Query("content", "")
	err := f.service.AddTodo(content)
	if err != nil {
		if errors.Is(err, InvalidContent) {
			return ctx.SendStatus(fiber.StatusBadRequest)
		} else {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}
	return ctx.SendStatus(fiber.StatusOK)
}
