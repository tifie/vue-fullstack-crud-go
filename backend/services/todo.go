package services

import (
	"go-todo/models"
	"go-todo/repositories"
	"go-todo/validators"

	"github.com/samber/do"
)

type ITodoService interface {
	GetTodos() (todos []models.TodoResponse, err error)
	CreateTodo(todo models.Todo) error
	DeleteTodo(todo models.Todo) error
}

type STodoService struct {
	todoRepository repositories.ITodoRepository
	todoValidator validators.ITodoValidator
}

func NewTodoService(i *do.Injector) (ITodoService, error) {
	repo := do.MustInvoke[repositories.ITodoRepository](i)
	vali := do.MustInvoke[validators.ITodoValidator](i)
	return &STodoService{todoRepository: repo, todoValidator: vali}, nil
}

func (ts *STodoService) GetTodos() (todos []models.TodoResponse, err error) {
	result, err := ts.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		todo := models.TodoResponse {
			ID: value.ID,
			Content: value.Content,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (ts *STodoService) CreateTodo(todo models.Todo) error {
	if err := ts.todoValidator.TodoValidate(todo); err != nil {
		return err
	}
	ts.todoRepository.CreateTodo(todo)
	return nil
}

func (ts *STodoService) DeleteTodo(todo models.Todo) error {
	ts.todoRepository.DeleteById(todo.ID)
	return nil
}