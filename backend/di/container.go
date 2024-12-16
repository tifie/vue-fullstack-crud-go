package di

import (
	"go-todo/controllers"
	"go-todo/repositories"
	"go-todo/services"

	"github.com/samber/do"
)

func NewContainer(injector *do.Injector) {
	do.Provide(injector, repositories.NewTodoRepository)
	do.Provide(injector, services.NewTodoService)
	do.Provide(injector, controllers.NewTodoController)
}