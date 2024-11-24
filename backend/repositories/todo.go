package repositories

import (
	"go-todo/models"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type ITodoRepository interface {
  FindAll() ([]models.Todo, error)
  DeleteById(todoId uint64) error
  CreateTodo(todo models.Todo) error
}

type todoRepository struct {
  db *gorm.DB
}

func NewTodoRepository(i *do.Injector) (ITodoRepository, error) {
  db := do.MustInvokeNamed[*gorm.DB](i, "sql")
  return &todoRepository{db}, nil
}

func (tr *todoRepository) FindAll() (todos []models.Todo, err error) {
  results := tr.db.Find(&todos)
  if results.Error != nil {
    return nil, results.Error
  }

  return todos, nil
}

func (tr *todoRepository) DeleteById(todoID uint64) error {
  var todo models.Todo
  result := tr.db.Where("ID = ?", todoID).Delete(&todo)
  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (tr *todoRepository) CreateTodo(todo models.Todo) error {
  if err := tr.db.Create(&models.Todo{Content: todo.Content}).Error; err != nil {
    return err
  }

  return nil
}