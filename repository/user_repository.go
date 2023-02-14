package repository

import "github.com/kleklai/todoAppv1/graph/model"

func (r *Repository) CreateUser(user model.User) (*model.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(id string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUser(id string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
