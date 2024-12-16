package services

import (
	"go-todo/models"
	"go-todo/repositories"

	"github.com/samber/do"
)

type ITodoService interface {
	FromList() (todos []models.TodoResponse, err error)
	CreateTodo(content string) error
	DeleteTodo(todo models.Todo) error
}

type STodoService struct {
	todoRepository repositories.ITodoRepository
}

func NewTodoService(i *do.Injector) (ITodoService, error) {
	// DIコンテナからレポジトリを取り出す
	repo := do.MustInvoke[repositories.ITodoRepository](i)
	return &STodoService{todoRepository: repo}, nil
}

func (ts *STodoService) FromList() (todos []models.TodoResponse, err error) {
	// データアクセス層からデータを取ってくる
	result, err := ts.todoRepository.List()
	// もし，エラーは発生したら，エラーを返す
	if err != nil {
		return nil, err
	}

	// データアクセス層から受け取ったデータの整形を行う
	for _, value := range result {
		todo := models.TodoResponse {
			ID: value.ID,
			Content: value.Content,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (ts *STodoService) CreateTodo(content string) error {
	todo := models.Todo{Content: content}
	return ts.todoRepository.Create(todo)
}

func (ts *STodoService) DeleteTodo(todo models.Todo) error {
	return ts.todoRepository.Delete(todo.ID)
}