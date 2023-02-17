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

	DB_USER := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE, DB_SSLMODE, DB_TIMEZONE)

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
	GetTodoByID(string) (*model.Todo, error)
	GetTodoByUser(string) ([]*model.Todo, error)
	GetTodoOfUserByStatus(string, bool) ([]*model.Todo, error)
	UpdateTodoDone(model.Todo) (*model.Todo, error)
	DeleteTodo(string) (*model.Todo, error)
}
