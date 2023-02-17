package service

import (
	"fmt"

	"github.com/kleklai/todoAppv1/graph/model"
)

func (s *Service) CreateTodo(todo *model.CreateTodoInput) (*model.Todo, error) {

	t := model.Todo{
		Text:   todo.Text,
		UserID: todo.UserID,
	}

	res, err := s.repoService.CreateTodo(t)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetTodoByID(id string) (*model.Todo, error) {

	if id == "" {
		return nil, fmt.Errorf("ID is empty")
	}

	res, err := s.repoService.GetTodoByID(id)

	if err != nil {
		return nil, fmt.Errorf("Todo not found")
	}

	return res, nil
}

func (s *Service) GetTodoByUser(userID string) ([]*model.Todo, error) {

	if userID == "" {
		return nil, fmt.Errorf("User ID is empty")
	}

	var todos []*model.Todo

	todos, err := s.repoService.GetTodoByUser(userID)

	if err != nil {
		return nil, fmt.Errorf("Todo not found")
	}

	return todos, nil
}

func (s *Service) GetTodoOfUserByStatus(userID string, done bool) ([]*model.Todo, error) {

	if userID == "" {
		return nil, fmt.Errorf("User ID is empty")
	}

	if _, err := s.repoService.GetUser(userID); err != nil {
		return nil, fmt.Errorf("User not found")
	}

	var todos []*model.Todo

	todos, err := s.repoService.GetTodoOfUserByStatus(userID, done)

	if err != nil {
		return nil, fmt.Errorf("Todo not found")
	}

	return todos, nil
}

func (s *Service) DeleteTodo(id string) (*model.Todo, error) {

	if id == "" {
		return nil, fmt.Errorf("ID is empty")
	}

	if _, err := s.repoService.GetTodoByID(id); err != nil {
		return nil, fmt.Errorf("Todo not found")
	}

	res, err := s.repoService.DeleteTodo(id)

	if err != nil {
		return nil, fmt.Errorf("Todo not found. Nothing to delete")
	}

	return res, nil
}

func (s *Service) UpdateTodoDone(todo *model.UpdateTodoDoneInput) (*model.UpdateTodoDone, error) {

	u := model.Todo{
		ID:   todo.ID,
		Done: todo.Done,
	}

	// Check if ID and DOne is empty
	if u.ID == "" {
		return nil, fmt.Errorf("ID is empty")
	}

	res, err := s.repoService.UpdateTodoDone(u)

	if err != nil {
		return nil, fmt.Errorf("Todo not found")
	}

	updateTodoDone := model.UpdateTodoDone{
		ID:   res.ID,
		Done: res.Done,
	}

	return &updateTodoDone, nil
}
