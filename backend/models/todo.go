package models

import (
	"time"
)

type Todo struct {
  ID        uint64 `gorm:"primarykey"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Content string `json:"content"`
}

type TodoResponse struct {
  ID uint64 `json:"id"`
  Content string `json:"content"`
}