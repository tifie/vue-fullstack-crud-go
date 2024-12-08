package main

import (
	"fmt"
	"go-todo/di"
	"go-todo/models"
	"go-todo/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
