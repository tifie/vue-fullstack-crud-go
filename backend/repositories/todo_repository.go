package repositories

import (
	"go-todo/models"

	"github.com/samber/do"
	"gorm.io/gorm"
)

// インターフェースの作成
// 継承っぽい実装が可能
type ITodoRepository interface {
  List() ([]models.Todo, error)
  Delete(todoId uint64) error
  Create(todo models.Todo) error
}

// 構造体の定義
type todoRepository struct {
  db *gorm.DB
}

func NewTodoRepository(i *do.Injector) (ITodoRepository, error) {
  // DIコンテナについては後日
  // DIコンテナからデータベースのコネクションプールを持ってくる
  db := do.MustInvokeNamed[*gorm.DB](i, "sql")
  return &todoRepository{db}, nil
}

func (tr *todoRepository) List() (todos []models.Todo, err error) {
  results := tr.db.Find(&todos)
  if results.Error != nil {
    return nil, results.Error
  }

  return todos, nil
}

func (tr *todoRepository) Delete(todoID uint64) error {
  var todo models.Todo
  result := tr.db.Where("ID = ?", todoID).Delete(&todo)
  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (tr *todoRepository) Create(todo models.Todo) error {
  if err := tr.db.Create(&models.Todo{Content: todo.Content}).Error; err != nil {
    return err
  }

  return nil
}