package service

import (
	"errors"
	"fmt"

	"github.com/kleklai/todoAppv1/graph/model"
)

func (s *Service) CreateUser(user *model.CreateUserInput) (*model.User, error) {

	_, err := s.repoService.GetUser(user.ID)

	// Validate if the use is already exist
	if err == nil {
		return nil, fmt.Errorf("User with id %s already exists", user.ID)
	}

	u := model.User{
		ID:   user.ID,
		Name: user.Name,
	}

	if u.ID == "" {
		// return nil, fmt.Errorf("ID is empty")
		return nil, errors.New("ID is empty")
	}

	res, err := s.repoService.CreateUser(u)

	if err != nil {
		return nil, fmt.Errorf("Error creating user: %v", err)
	}

	return res, nil
}

func (s *Service) GetUser(id string) (*model.User, error) {

	if id == "" {
		// return nil, fmt.Errorf("ID is empty")
		return nil, errors.New("ID is empty")
	}

	res, err := s.repoService.GetUser(id)

	if err != nil {
		return nil, errors.New("User not found")
	}

	return res, nil
}

func (s *Service) DeleteUser(id string) (*model.User, error) {

	if id == "" {
		return nil, errors.New("ID is empty")
	}

	if _, err := s.repoService.GetUser(id); err != nil {
		return nil, errors.New("User not found")
	}

	res, err := s.repoService.DeleteUser(id)

	if err != nil {
		return nil, errors.New("Error deleting user")
	}

	return res, nil
}
