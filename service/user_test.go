package service

import (
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/kleklai/todoAppv1/graph/model"
	"github.com/kleklai/todoAppv1/repository"
	"github.com/stretchr/testify/assert"
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

func TestCreateUser(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	user := model.CreateUserInput{
		ID:   generateRandomString(),
		Name: "Maynard",
	}

	t.Run("success", func(t *testing.T) {
		m.On("CreateUser", user).Return(&model.User{
			ID: user.ID,
		}, nil).Once()

		res, err := service.CreateUser(&user)

		assert.NotNil(t, res)
		assert.NoError(t, err)
		m.CreateUser(user)

		m.AssertCalled(t, "CreateUser", user)
		m.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	id := "5ycOZ54lCpPPuDdL-iGzLfwuHTbX"

	m.On("GetUser", id).Return(&model.User{}, nil).Once()

	res, err := service.GetUser(id)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.GetUser(id)

	m.AssertCalled(t, "GetUser", id)
	m.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {

	m := &RepositoryMock{}

	service := NewService(*repository.NewRepository())

	id := "5ycOZ54lCpPPuDdL-iGzLfwuHTbX"

	m.On("DeleteUser", id).Return(&model.User{}, nil).Once()

	res, err := service.DeleteUser(id)

	assert.NotNil(t, res)
	assert.NoError(t, err)

	m.DeleteUser(id)

	m.AssertCalled(t, "DeleteUser", id)
	m.AssertExpectations(t)
}
