package validators

import (
	"errors"
	"go-todo/models"

	"github.com/samber/do"
)

type ITodoValidator interface {
	TodoValidate(todo models.Todo) error
}

type STodoValidator struct {}

func NewTodoValidator(i *do.Injector) (ITodoValidator, error) {
	return &STodoValidator{}, nil
}

func (sv *STodoValidator) TodoValidate(todo models.Todo) error {
	if todo.Content != "" {
		return errors.New("Content can't be blank")
	}
	return nil
}