package models

import (
	"time"
)

/*
	  モデルの宣言の仕方は以下のURLに技術してある
		https://gorm.io/ja_JP/docs/models.html
*/
type Todo struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content" binding:"required"`
}

type TodoResponse struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
}
