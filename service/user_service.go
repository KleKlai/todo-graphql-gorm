package service

import "github.com/kleklai/todoAppv1/graph/model"

func (s *Service) CreateUser(user *model.CreateUserInput) (*model.User, error) {

	u := model.User{
		ID:   user.ID,
		Name: user.Name,
	}

	// if user.ID == "" {
	// 	return nil, error("ID is required")
	// }

	res, err := s.repoService.CreateUser(u)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetUser(id string) (*model.User, error) {

	res, err := s.repoService.GetUser(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteUser(id string) (*model.User, error) {

	res, err := s.repoService.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
