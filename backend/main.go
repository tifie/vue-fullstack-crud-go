package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// dbとのデータのやり取りをするモデルの作成
type Todo struct {
	ID        uint64    `gorn:"primarykey"`
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
		//すべてのレコードを取得
		result := DB.Find(&todos)
		if result.Error != nil {
			log.Fatalf("Error GEt todos list: %v", result.Error)
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

		//INSERT INTO todo VALUE ("content");
		result := DB.Create(&Todo{Content: content})
		if result.Error != nil {
			log.Fatalf("Error Create todo: %v", result.Error)
			return
		}
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.DELETE("/todo/delete", func(ctx *gin.Context) {
		ids := ctx.PostForm("id")
		id, _ := strconv.ParseUint(ids, 10, 64) //10進，64bit
		// DB.Delete(&Todo{ID: id})
		DB.Delete(&Todo{}, id) //宿題修正点
	})

	r.PUT("/todo/edit", func(ctx *gin.Context) { //宿題修正点
		content := ctx.PostForm("content")
		fmt.Println(ctx.Request.PostForm, content)
		if content == "" {
			return
		}
		ids := ctx.PostForm("id")
		id, _ := strconv.ParseUint(ids, 10, 64) //10進，64bit
		DB.Model(&Todo{}).Where("id = ?", id).Update("Content", content)
	})
}

//ここからmain関数

func main() {
	r := gin.Default()

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
	DB = db

	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatalf("Failed to Migrate Todo scheem")
	}

	//router の設定
	listners(r)

	r.Run()
}
