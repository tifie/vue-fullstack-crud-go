package main

import (
	"fmt"
	"go-todo/di"
	"go-todo/models"
	"go-todo/routers"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Todo struct {
	ID        uint64    `gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
}

func listners(r *gin.Engine) {
  r.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.GET("/todo/list", func(ctx *gin.Context) {
    var todos []Todo
		result := DB.Find(&todos)
    if result.Error != nil {
      log.Fatalf("Error Get todos list: %v", result.Error)
      return
    }

    ctx.JSON(http.StatusOK, todos)
  })

  r.POST("/todo/create", func(ctx *gin.Context) {
    content := ctx.PostForm("content")
    fmt.Println(ctx.Request.PostForm, content)
    if content == "" {
      return
    }
    result := DB.Create(&Todo{Content: content}) // INSERT INTO todo VALUE("content");
    if result.Error != nil {
      log.Fatalf("Error Create todos: %v", result.Error)
      return
    }
    ctx.Redirect(http.StatusMovedPermanently, "/index")
  })

  r.DELETE("/todo/delete", func(ctx *gin.Context) {
    ids := ctx.Query("id")
    id, _ := strconv.ParseUint(ids, 10, 64)
    DB.Delete(&Todo{ID: id})
  })

  r.PUT("/todo/edit", func(ctx *gin.Context) {
		ids := ctx.Query("id")
		content := ctx.Query("content")
		id, _ := strconv.ParseUint(ids, 10, 64)

		DB.Model(&Todo{}).Where("ID = ?", id).Update("content", content)
	})
}

func main() {
  r := gin.Default()

  // 環境変数を読み込む
  err := godotenv.Load()
  if err != nil {
		log.Fatal(".env file failed to load!!")
  }

  dsn := fmt.Sprintf(
    "host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_DATABASE"),
    os.Getenv("DB_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
    log.Fatalf("Falied to connect to database: %v", err)
  }
  err = db.AutoMigrate(&models.Todo{})
  if err != nil {
    log.Fatalf("Falied to migrate databse: %v", err)
  }

  injector := do.New()
  do.ProvideNamed[*gorm.DB](injector, "sql", func(i *do.Injector) (*gorm.DB, error) {
    return db, nil
  })
  di.NewContainer(injector)

  routers.SetupRouter(r, injector)

  r.Run()
}