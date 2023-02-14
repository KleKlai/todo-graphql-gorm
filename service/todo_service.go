package service

import "github.com/kleklai/todoAppv1/graph/model"

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

func (s *Service) GetTodoByUser(userID string) ([]*model.Todo, error) {

	var todos []*model.Todo

	todos, err := s.repoService.GetTodoByUser(userID)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Service) DeleteTodo(id string) (*model.Todo, error) {

	res, err := s.repoService.DeleteTodo(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
