package service

import (
	"testing"

	"github.com/kleklai/todoAppv1/graph/model"
	repository "github.com/kleklai/todoAppv1/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

func (m *RepositoryMock) UpdateTodoDone(todo *model.UpdateTodoDoneInput) (*model.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) UpdateTodoTask(todo *model.UpdateTodoTaskInput) (*model.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *RepositoryMock) DeleteTodo(id string) error {
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

	m.On("CreateTodo", mock.Anything).Return(&model.Todo{
		Task: todo.Task,
	}, nil)

	_, err := service.CreateTodo(&todo)

	// assert.NotNil(t, res)
	assert.NoError(t, err)
	require.Nil(t, err)

	// m.CreateTodo(todo)

	// m.AssertCalled(t, "CreateTodo", todo)
	m.AssertExpectations(t)
}

func TestGetTodoByID(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	id := "a1f04dff-feb3-44d1-891d-3184c6f4c18f"

	m.On("GetTodoByID", id).Return(&model.Todo{}, nil).Once()

	res, err := service.GetTodoByID(id)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.GetTodoByID(id)

	m.AssertCalled(t, "GetTodoByID", id)
	m.AssertExpectations(t)
}

func TestGetTodoByUserID(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	userID := "6ikyYWd9iOwRTjBUu5LFws2cRAy7"

	m.On("GetTodoByUserID", userID).Return([]*model.Todo{}, nil).Once()

	res, err := service.GetTodoByUserID(userID)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.GetTodoByUserID(userID)

	m.AssertCalled(t, "GetTodoByUserID", userID)
	m.AssertExpectations(t)
}

func TestGetTodoOfUserByStatus(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	userID := "ZazIuf_YELWeYQoHcBJjcJ5kkC3m"
	status := true

	m.On("GetTodoOfUserByStatus", userID, status).Return([]*model.Todo{}, nil).Once()

	res, err := service.GetTodoOfUserByStatus(userID, status)

	assert.NotNil(t, res)
	assert.NoError(t, err)
	m.GetTodoOfUserByStatus(userID, status)

	m.AssertCalled(t, "GetTodoOfUserByStatus", userID, status)
	m.AssertExpectations(t)
}

func TestUpdateTodoDone(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	input := model.UpdateTodoDoneInput{
		ID:   "a7d337a1-871d-4605-9f01-2e531cb8a790",
		Done: true,
	}

	m.On("UpdateTodoDone", &input).Return(&model.Todo{}, nil).Once()

	res, err := service.UpdateTodoDone(&input)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.UpdateTodoDone(&input)

	m.AssertCalled(t, "UpdateTodoDone", &input)
	m.AssertExpectations(t)
}

func TestUpdateTodoTask(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	input := model.UpdateTodoTaskInput{
		ID:   "a7d337a1-871d-4605-9f01-2e531cb8a790",
		Task: "Test Todo",
	}

	m.On("UpdateTodoTask", &input).Return(&model.Todo{}, nil).Once()

	res, err := service.UpdateTodoTask(&input)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.UpdateTodoTask(&input)

	m.AssertCalled(t, "UpdateTodoTask", &input)
	m.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	id := "635d9c7f-9cc0-42da-a231-6e56973c723c"

	m.On("DeleteTodo", id).Return(nil).Once()

	res, err := service.DeleteTodo(id)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.DeleteTodo(id)

	m.AssertCalled(t, "DeleteTodo", id)
	m.AssertExpectations(t)
}
