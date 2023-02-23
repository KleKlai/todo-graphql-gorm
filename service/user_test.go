package service

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/kleklai/todoAppv1/graph/model"
	"github.com/stretchr/testify/mock"
)

func generateRandomString() string {
	// Generate random 28char string
	randBytes := make([]byte, 32)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}

	// Convert the random bytes to a base64 string
	randString := base64.RawURLEncoding.EncodeToString(randBytes)

	// Trim the string to 28 characters
	randString = randString[:28]
	return randString
}

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) CreateUser(input model.CreateUserInput) (*model.User, error) {

	args := m.Called(input)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *RepositoryMock) GetUser(id string) (*model.User, error) {

	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *RepositoryMock) DeleteUser(id string) (*model.User, error) {

	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *RepositoryMock) CreateTodo(input model.CreateTodoInput) (*model.Todo, error) {

	args := m.Called(input)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoByID(id string) (*model.Todo, error) {

	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoByUserID(id string) ([]*model.Todo, error) {

	args := m.Called(id)
	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoOfUserByStatus(id string, done bool) ([]*model.Todo, error) {

	args := m.Called(id, done)
	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (m *RepositoryMock) UpdateTodoDone(todo model.Todo) (*model.Todo, error) {

	args := m.Called(todo)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) UpdateTodoTask(todo model.Todo) (*model.Todo, error) {

	args := m.Called(todo)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) DeleteTodo(id string) (*model.Todo, error) {

	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func TestCreateUser(t *testing.T) {

	mock := &RepositoryMock{}
	service := NewService(mock)

	// input := model.CreateUserInput{
	// 	ID:   generateRandomString(),
	// 	Name: "Maynard",
	// }
}

// func TestCreateUser(t *testing.T) {

// 	// Open Database Connection
// 	service := NewService(*repository.NewRepository())

// 	testInput := &model.CreateUserInput{
// 		ID:   generateRandomString(),
// 		Name: "Maynard",
// 	}

// 	res, err := service.CreateUser(testInput)

// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	if res.ID != testInput.ID {
// 		t.Errorf("Expected id to be %s, but got %s", testInput.ID, res.ID)
// 	}

// 	if res.Name != testInput.Name {
// 		t.Errorf("Expected name to be %s, but got %s", testInput.Name, res.Name)
// 	}

// 	dbTodo, err := service.GetUser(res.ID)

// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if dbTodo.ID != testInput.ID {
// 		t.Errorf("Expected id to be %s, but got %s", testInput.ID, dbTodo.ID)
// 	}

// 	if dbTodo.Name != testInput.Name {
// 		t.Errorf("Expected nameto be %s, but got %s", testInput.Name, dbTodo.Name)
// 	}
// }

// func TestGetUser(t *testing.T) {

// 	// Open Database Connection
// 	service := NewService(*repository.NewRepository())

// 	testInput := &model.CreateUserInput{
// 		ID:   generateRandomString(),
// 		Name: "Maynard",
// 	}

// 	_, err := service.GetUser(testInput.ID)

// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	dbTodo, err := service.GetUser(testInput.ID)

// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if dbTodo.ID != testInput.ID {
// 		t.Errorf("Expected id to be %s, but got %s", testInput.ID, dbTodo.ID)
// 	}

// 	if dbTodo.Name != testInput.Name {
// 		t.Errorf("Expected nameto be %s, but got %s", testInput.Name, dbTodo.Name)
// 	}
// }
