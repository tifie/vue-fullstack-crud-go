package routers

import (
	"go-todo/controllers"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func todoRouter(router *gin.Engine, injector *do.Injector) {
	todoHandler := do.MustInvoke[controllers.ITodoController](injector)
	router.POST("/todo/create", todoHandler.CreateTodo)
	router.GET("/todo/lists", todoHandler.FromList)
	router.DELETE("/todo/delete/:id", todoHandler.DeleteTodo)
}