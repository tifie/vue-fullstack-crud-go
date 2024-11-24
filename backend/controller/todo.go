package controller

import (
	"fmt"
	"go-todo/models"
	"go-todo/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ITodoController interface {
	PostTodo(ctx *gin.Context)
	GetTodos(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type STodoController struct {
	todoService services.ITodoService
}

func NewTodoController(i *do.Injector) (ITodoController, error) {
	service := do.MustInvoke[services.ITodoService](i)
	return &STodoController{todoService: service}, nil
}

func (tc *STodoController) PostTodo(ctx *gin.Context) {
	var todo models.Todo
	todo.Content = ctx.PostForm("content")
	fmt.Println(ctx.Request.PostForm, todo.Content)
	fmt.Println(todo)
	if err := tc.todoService.CreateTodo(todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/index")
}

func (tc *STodoController) GetTodos(ctx *gin.Context){
	todos, err := tc.todoService.GetTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (tc *STodoController) DeleteTodo(ctx *gin.Context){
	ids := ctx.PostForm("id")
	id, _ := strconv.ParseUint(ids, 10, 64)
	todo := models.Todo{ID: id}
	tc.todoService.DeleteTodo(todo)
}