package main

import (
	"fmt"
	"go-todo/models"
	"go-todo/routers"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func listeners(r *gin.Engine) {
  r.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.POST("/todo/create", func(ctx *gin.Context) {
    content := ctx.PostForm("content")
    fmt.Println(ctx.Request.PostForm, content)
    if content == "" {
      return
    }
    result := models.DB.Create(&models.Todo{Content: content}) // INSERT INTO todo VALUE("content");
    if result.Error != nil {
      log.Fatalf("Error Create todos: %v", result.Error)
      return
    }
    ctx.Redirect(http.StatusMovedPermanently, "/index")
  })

  r.GET("/todo/list", func(ctx *gin.Context) {
    var todos []models.Todo
    // 全てのレコードを取得
    result := models.DB.Find(&todos) // SELECT * FROM todo;
    if result.Error != nil {
      log.Fatalf("Error Get todos list: %v", result.Error)
      return
    }
    // fmt.Println(json.NewEncoder(os.Stdout).Encode(todos))
    ctx.JSON(http.StatusOK, todos)
  })

  r.DELETE("/todo/delete", func(ctx *gin.Context) {
    ids := ctx.PostForm("id")
    id, _ := strconv.ParseUint(ids, 10, 64)
    models.DB.Delete(&models.Todo{ID: id})
  })
}

func main() {
  r := gin.Default()

  // 環境変数を読み込む
  err := godotenv.Load()
  if err != nil {
    log.Fatal(".env file failed to load!")
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

  injector := do.New()
  do.ProvideNamed[*gorm.DB](injector, "sql", func(i *do.Injector) (*gorm.DB, error) {
    return db, nil
  })

  if err != nil {
    log.Fatalf("Falied to migrate databse: %v", err)
  }

  routers.SetupRouter(r, injector)
  // listeners(r)

  r.Run()
}
