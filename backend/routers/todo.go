package routers

import (
	"go-todo/controller"
	"go-todo/repositories"
	"go-todo/services"
	"go-todo/validators"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func TodoRouter(router *gin.Engine, injector *do.Injector) {
	do.Provide(injector, repositories.NewTodoRepository)
	do.Provide(injector, services.NewTodoService)
	do.Provide(injector, controller.NewTodoController)
	do.Provide(injector, validators.NewTodoValidator)
	todoHandler := do.MustInvoke[controller.ITodoController](injector)
	router.POST("/todo/create", todoHandler.PostTodo)
	router.GET("/todo/lists", todoHandler.GetTodos)
	router.DELETE("/todo/delete", todoHandler.DeleteTodo)
}