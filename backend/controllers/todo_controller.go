package controller

import (
	"go-todo/models"
	"go-todo/services"
	"go-todo/validators"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type ITodoController interface {
	CreateTodo(ctx *gin.Context)
	FromList(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	EditTodo(ctx *gin.Context)
}

type STodoController struct {
	todoService services.ITodoService
}

func NewTodoController(i *do.Injector) (ITodoController, error) {
	service := do.MustInvoke[services.ITodoService](i)
	return &STodoController{todoService: service}, nil
}

func (tc *STodoController) CreateTodo(ctx *gin.Context) {
	var todo = models.Todo{}

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	if !validators.CreateTodoValidation(todo, ctx) {
		return
	}

	if err := tc.todoService.Create(todo); err != nil {
		log.Fatalf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/index")
}

func (tc *STodoController) EditTodo(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	todo := models.Todo{ID: id}
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	if !validators.CreateTodoValidation(todo, ctx) {
		return
	}
	tc.todoService.Edit(todo)
}

func (tc *STodoController) FromList(ctx *gin.Context) {
	todos, err := tc.todoService.List()
	if err != nil {
		log.Fatalf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (tc *STodoController) DeleteTodo(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	todo := models.Todo{ID: id}
	tc.todoService.Delete(todo)
}
