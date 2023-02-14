package repository

import "github.com/kleklai/todoAppv1/graph/model"

func (r *Repository) CreateTodo(todo model.Todo) (*model.Todo, error) {

	if err := r.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *Repository) GetTodoByUser(userID string) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := r.db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *Repository) DeleteTodo(id string) (*model.Todo, error) {
	var todo model.Todo

	if err := r.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
