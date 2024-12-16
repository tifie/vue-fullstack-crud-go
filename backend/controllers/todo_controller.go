package controllers

import (
	"go-todo/models"
	"go-todo/services"

	//"go-todo/validators"
	"go-todo/validations"

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

  //if !validators.CreateTodoValidation(todo, ctx) {
	if !validations.CreateTodoValidation(todo, ctx) {
		return
	}

	if err := tc.todoService.CreateTodo(todo.Content); err != nil {
		log.Fatalf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/index")
}

func (tc *STodoController) FromList(ctx *gin.Context){
	todos, err := tc.todoService.FromList()
	if err != nil {
		log.Fatalf("Error: %v", err)
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