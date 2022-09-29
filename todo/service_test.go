package todo_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mocks "yusa/mocks/todo"
	"yusa/todo"
)

func TestNewService(t *testing.T) {
	t.Run("given valid repository then it should return service when NewService called", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepository := mocks.NewMockTodoRepository(ctrl)
		service := todo.NewService(mockRepository)
		assert.NotNil(t, service)
	})
	t.Run("given not valid repository then it should return nil when NewService called", func(t *testing.T) {
		service := todo.NewService(nil)
		assert.Nil(t, service)
	})
}

func TestService_AddTodo(t *testing.T) {
	t.Run("given empty content then it should return error when AddTodo called", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepository := mocks.NewMockTodoRepository(ctrl)
		service := todo.NewService(mockRepository)
		err := service.AddTodo("")
		assert.NotNil(t, err)
	})
	t.Run("given not valid content then it should return error when AddTodo called", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepository := mocks.NewMockTodoRepository(ctrl)
		service := todo.NewService(mockRepository)
		err := service.AddTodo("sd")
		assert.NotNil(t, err)
	})
	t.Run("given already exist content then it should return error when AddTodo called", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepository := mocks.NewMockTodoRepository(ctrl)
		service := todo.NewService(mockRepository)
		content := "sdasfsaf"
		mockRepository.EXPECT().HasTodo(content).Return(true, nil).MinTimes(1)
		err := service.AddTodo(content)

		assert.NotNil(t, err)
	})
	t.Run("given valid and not exist content then it return nil when AddTodo called", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepository := mocks.NewMockTodoRepository(ctrl)
		service := todo.NewService(mockRepository)
		content := "sdasfsaf"
		mockRepository.EXPECT().HasTodo(content).Return(false, nil).MinTimes(1)
		mockRepository.EXPECT().AddTodo(gomock.Any()).Return(nil).MinTimes(1)
		err := service.AddTodo(content)
		assert.Nil(t, err)
	})
}
