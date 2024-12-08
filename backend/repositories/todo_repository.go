package repositories

import (
	"go-todo/models"

	"github.com/samber/do"
	"gorm.io/gorm"
)

// インターフェースの作成
// 継承っぽい実装が可能
type ITodoRepository interface {
	FindAll() ([]models.Todo, error)
	DeleteBy(todoId uint64) error
	EditBy(todo models.Todo) error
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

func (tr *todoRepository) FindAll() (todos []models.Todo, err error) {
	result := tr.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (tr *todoRepository) DeleteBy(todoID uint64) error {
	var todo models.Todo
	result := tr.db.Where("ID = ?", todoID).Delete(&todo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tr *todoRepository) EditBy(todo models.Todo) error {
	result := tr.db.Model(&models.Todo{}).Where("ID = ?", todo.ID).Update("content", todo.Content)
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
