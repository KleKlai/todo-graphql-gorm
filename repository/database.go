package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kleklai/todoAppv1/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect_db() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// source := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_DATABASE") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), "disable", "Asia/Shanghai")

	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

	if err != nil {
		fmt.Println(source)
		log.Fatal(err)
		panic("failed to connect database")
	}

	return db
}

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{
		db: connect_db(),
	}
}

type RepositoryInterface interface {
	CreateUser(model.User) (*model.User, error)
	GetUser(string) (*model.User, error)
	DeleteUser(string) (*model.User, error)

	CreateTodo(model.Todo) (*model.Todo, error)
	GetTodoByUser(string) ([]*model.Todo, error)
	DeleteTodo(string) (*model.Todo, error)
}
