package service

import (
	"errors"

	"github.com/kleklai/todoAppv1/graph/model"
)

func (s *Service) CreateTodo(todo *model.CreateTodoInput) (*model.Todo, error) {

	t := model.Todo{
		Task:   todo.Task,
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
		return nil, errors.New("ID is empty")
	}

	res, err := s.repoService.GetTodoByID(id)

	if err != nil {
		return nil, errors.New("Todo not found")
	}

	return res, nil
}

func (s *Service) GetTodoByUserID(userID string) ([]*model.Todo, error) {

	if userID == "" {
		return nil, errors.New("User ID is empty")
	}

	var todos []*model.Todo

	todos, err := s.repoService.GetTodoByUserID(userID)

	if err != nil {
		return nil, errors.New("Todo not found")
	}

	return todos, nil
}

func (s *Service) GetTodoOfUserByStatus(userID string, done bool) ([]*model.Todo, error) {

	if userID == "" {
		return nil, errors.New("User ID is empty")
	}

	if _, err := s.repoService.GetUser(userID); err != nil {
		return nil, errors.New("User not found")
	}

	var todos []*model.Todo

	todos, err := s.repoService.GetTodoOfUserByStatus(userID, done)

	if err != nil {
		return nil, errors.New("Todo not found")
	}

	return todos, nil
}

func (s *Service) DeleteTodo(id string) (*model.Todo, error) {

	if id == "" {
		return nil, errors.New("ID is empty")
	}

	if _, err := s.repoService.GetTodoByID(id); err != nil {
		return nil, errors.New("Todo not found")
	}

	res, err := s.repoService.DeleteTodo(id)

	if err != nil {
		return nil, errors.New("Todo not found")
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
		return nil, errors.New("ID is empty")
	}

	res, err := s.repoService.UpdateTodoDone(u)

	if err != nil {
		return nil, errors.New("Todo not found")
	}

	updateTodoDone := model.UpdateTodoDone{
		ID:   res.ID,
		Done: res.Done,
	}

	return &updateTodoDone, nil
}

func (s *Service) UpdateTodoTask(todo *model.UpdateTodoTaskInput) (*model.UpdateTodoTask, error) {

	u := model.Todo{
		ID:   todo.ID,
		Task: todo.Task,
	}

	// Check if ID and Task is empty
	if u.ID == "" {
		return nil, errors.New("ID is empty")
	}

	res, err := s.repoService.UpdateTodoTask(u)

	if err != nil {
		return nil, errors.New("Todo not found")
	}

	updateTodoTask := model.UpdateTodoTask{
		ID:   res.ID,
		Task: res.Task,
	}

	return &updateTodoTask, nil
}
