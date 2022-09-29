package todo_test

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	mocks "yusa/mocks/todo"
	"yusa/todo"
)

func TestFiberHandler_AddTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mocks.NewMockTodoService(ctrl)
	handler := todo.NewHandler(mockService)
	t.Run("given valid todo content then it should return statuk ok", func(t *testing.T) {
		content := "asdafas"
		req := httptest.NewRequest("POST", "/todos/?content="+content, nil)
		app := fiber.New()
		mockService.EXPECT().AddTodo(content).Return(nil).MinTimes(1)
		app.Post("/todos", handler.AddTodo)
		resp, _ := app.Test(req)
		assert.Equal(t, resp.StatusCode, fiber.StatusOK)
	})
	t.Run("given invalid content then it should return InvalidContent error.", func(t *testing.T) {
		content := ""
		req := httptest.NewRequest("POST", "/todos/?content="+content, nil)
		app := fiber.New()
		mockService.EXPECT().AddTodo(content).Return(todo.InvalidContent).MinTimes(1)
		app.Post("/todos", handler.AddTodo)
		resp, _ := app.Test(req)
		assert.Equal(t, resp.StatusCode, fiber.StatusBadRequest)
	})
	t.Run("given invalid content then it should return InvalidContent error.", func(t *testing.T) {
		content := "asdad"
		req := httptest.NewRequest("POST", "/todos/?content="+content, nil)
		app := fiber.New()
		mockService.EXPECT().AddTodo(content).Return(errors.New("test error")).MinTimes(1)
		app.Post("/todos", handler.AddTodo)
		resp, _ := app.Test(req)
		assert.Equal(t, resp.StatusCode, fiber.StatusInternalServerError)
	})
}
