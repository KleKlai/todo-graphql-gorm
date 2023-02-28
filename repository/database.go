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

	var err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		dbUser     = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbDatabase = os.Getenv("DB_DATABASE")
		dbSsl      = os.Getenv("DB_SSLMODE")
	)

	// var (
	// 	dbUser     = "postgres"
	// 	dbPassword = "postgres"
	// 	dbHost     = "localhost"
	// 	dbPort     = "5432"
	// 	dbDatabase = "todo"
	// 	dbSsl      = "disable"
	// )

	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbSsl, "Asia/Singapore")

	db, err := gorm.Open(postgres.Open(source), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)

		return nil
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
