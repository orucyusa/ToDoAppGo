package todo

import (
	"errors"
	"time"
)

type Todo struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	InvalidContent     = errors.New("given content is invalid to add as todo")
	AlreadyFindContent = errors.New("already has that content")
)

const (
	MinimumTodoContentLength = 5
	MaximumTodoContentLength = 50
)
