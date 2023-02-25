package repository

import (
	"fmt"
	"log"

	"github.com/kleklai/todoAppv1/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect_db() *gorm.DB {

	// err := godotenv.Load()

	// DB_USER := os.Getenv("DB_USERNAME")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	// DB_DATABASE := os.Getenv("DB_DATABASE")
	// DB_SSLMODE := os.Getenv("DB_SSLMODE")
	// DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	DB_USER := "postgres"
	DB_PASSWORD := "postgres"
	DB_HOST := "localhost"
	DB_PORT := "5432"
	DB_DATABASE := "todo"
	DB_SSLMODE := "disable"
	DB_TIMEZONE := "Asia/Singapore"

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

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
	GetTodoByUserID(string) ([]*model.Todo, error)
	GetTodoOfUserByStatus(string, bool) ([]*model.Todo, error)
	UpdateTodoDone(model.Todo) (*model.Todo, error)
	UpdateTodoTask(model.Todo) (*model.Todo, error)
	DeleteTodo(string) (*model.Todo, error)
}
