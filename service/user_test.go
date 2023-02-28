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
		ID:   "7fdde9ab-a814-4738-bfe5-79e38245dafa",
		Name: "Maynard",
	}

	m.On("GetUser", mock.Anything).Return(&model.User{
		ID: "7fdde9ab-a814-4738-bfe5-79e38245dafa",
	}, nil)

	m.On("CreateUser", mock.Anything).Return(&model.User{
		ID: "7fdde9ab-a814-4738-bfe5-79e38245dafa",
	}, nil)

	// m.On("CreateUser", mock.Anything).Return(&model.User{
	// 	ID: "7fdde9ab-a814-4738-bfe5-79e38245dafa",
	// }, nil)

	_, err := service.CreateUser(&user)

	// assert.NotNil(t, res)
	assert.NoError(t, err)
	// m.CreateUser(user)

	// m.AssertCalled(t, "CreateUser", user)
	// m.AssertExpectations(t)

	// t.Run("GetUser", func(t *testing.T) {
	// 	m.On("GetUser", res.ID).Return(&model.User{}, nil).Once()

	// 	res, err := service.GetUser(res.ID)

	// 	assert.NotNil(t, res)
	// 	assert.NoError(t, err)

	// 	m.GetUser(res.ID)

	// 	m.AssertCalled(t, "GetUser", res.ID)
	// 	m.AssertExpectations(t)
	// })
}

// func TestGetUser(t *testing.T) {

// 	m := &RepositoryMock{}

// 	service := NewService(*repository.NewRepository())

// 	id := "6ikyYWd9iOwRTjBUu5LFws2cRAy7"

// 	m.On("GetUser", id).Return(&model.User{}, nil).Once()

// 	res, err := service.GetUser(id)

// 	assert.NotNil(t, res)
// 	assert.NoError(t, err)

// 	m.GetUser(id)

// 	m.AssertCalled(t, "GetUser", id)
// 	m.AssertExpectations(t)
// }

// func TestDeleteUser(t *testing.T) {

// 	m := &RepositoryMock{}

// 	service := NewService(*repository.NewRepository())

// 	id := "7LPh_DQHGiH7JCCrCPd5Mgn8TkhL"

// 	m.On("DeleteUser", id).Return(&model.User{}, nil).Once()

// 	res, err := service.DeleteUser(id)

// 	assert.NotNil(t, res)
// 	assert.NoError(t, err)

// 	m.DeleteUser(id)

// 	m.AssertCalled(t, "DeleteUser", id)
// 	m.AssertExpectations(t)
// }
