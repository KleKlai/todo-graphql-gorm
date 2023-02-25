package service

import (
	"testing"

	"github.com/kleklai/todoAppv1/graph/model"
	repository "github.com/kleklai/todoAppv1/repository"
	"github.com/stretchr/testify/assert"
)

func (m *RepositoryMock) CreateTodo(todo model.CreateTodoInput) (*model.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoByID(id string) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoByUserID(userID string) ([]*model.Todo, error) {
	args := m.Called(userID)
	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (m *RepositoryMock) GetTodoOfUserByStatus(userID string, done bool) ([]*model.Todo, error) {
	args := m.Called(userID, done)
	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (m *RepositoryMock) UpdateTodoStatus(id string, done bool) (*model.Todo, error) {
	args := m.Called(id, done)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) UpdateTodoTask(id string, task string) (*model.Todo, error) {
	args := m.Called(id, task)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) DeleteTodoByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTodo(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	todo := model.CreateTodoInput{
		Task:   "Test Todo",
		UserID: "6ikyYWd9iOwRTjBUu5LFws2cRAy7",
	}

	m.On("CreateTodo", todo).Return(&model.Todo{
		Task: todo.Task,
	}, nil).Once()

	res, err := service.CreateTodo(&todo)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.CreateTodo(todo)

	m.AssertCalled(t, "CreateTodo", todo)
	m.AssertExpectations(t)
}

func TestGetTodoByID(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	id := "a7d337a1-871d-4605-9f01-2e531cb8a790"

	m.On("GetTodoByID", id).Return(&model.Todo{}, nil).Once()

	res, err := service.GetTodoByID(id)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.GetTodoByID(id)

	m.AssertCalled(t, "GetTodoByID", id)
	m.AssertExpectations(t)
}
