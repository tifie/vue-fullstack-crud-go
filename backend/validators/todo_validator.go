package validators

import (
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoValidation(todo models.Todo, ctx *gin.Context) bool {

	// 文字制限を付けたい場合
	if len(todo.Content) > 128 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content must be at least 128 characters long"})
		return false
	}

	return true
}
