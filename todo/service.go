package todo

import "github.com/rs/xid"

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	AddTodo(todo Todo) error
	HasTodo(content string) (bool, error)
}
type Service struct {
	repo TodoRepository
}

func NewService(repo TodoRepository) *Service {
	if repo == nil {
		return nil
	}
	return &Service{repo: repo}
}
func (s *Service) AddTodo(content string) error {
	if content == "" {
		return InvalidContent
	}
	if len(content) < MinimumTodoContentLength || len(content) > MaximumTodoContentLength {
		return InvalidContent
	}
	isAlreadyExistContent, err := s.repo.HasTodo(content)
	if err != nil {
		return err
	}
	if isAlreadyExistContent {
		return AlreadyFindContent
	}

	id := xid.New().String()
	err = s.repo.AddTodo(Todo{
		ID:      id,
		Content: content,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTodos() ([]Todo, error) {
	return s.repo.GetTodos()
}
