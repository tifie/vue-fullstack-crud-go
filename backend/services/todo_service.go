package services

import (
	"go-todo/models"
	"go-todo/repositories"

	"github.com/samber/do"
)

type ITodoService interface {
	List() (todos []models.TodoResponse, err error)
	Create(todo models.Todo) error
	Delete(todo models.Todo) error
	Edit(todo models.Todo) error
}

type STodoService struct {
	todoRepository repositories.ITodoRepository
}

func NewTodoService(i *do.Injector) (ITodoService, error) {
	// DIコンテナからレポジトリを取り出す
	repo := do.MustInvoke[repositories.ITodoRepository](i)
	return &STodoService{todoRepository: repo}, nil
}

func (ts *STodoService) List() (todos []models.TodoResponse, err error) {
	// データアクセス層からデータを取ってくる
	result, err := ts.todoRepository.FindAll()

	// もし，エラーは発生したら，エラーを返す
	if err != nil {
		return nil, err
	}

	// データアクセス層から受け取ったデータの整形を行う
	for _, value := range result {
		todo := models.TodoResponse{
			ID:      value.ID,
			Content: value.Content,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (ts *STodoService) Create(todo models.Todo) error {
	return ts.todoRepository.Create(todo)
}

func (ts *STodoService) Delete(todo models.Todo) error {
	return ts.todoRepository.DeleteBy(todo.ID)
}

func (ts *STodoService) Edit(todo models.Todo) error {
	return ts.todoRepository.EditBy(todo)
}
