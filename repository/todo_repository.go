package repository

import "github.com/kleklai/todoAppv1/graph/model"

func (r *Repository) CreateTodo(todo model.Todo) (*model.Todo, error) {

	if err := r.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *Repository) GetTodoByID(id string) (*model.Todo, error) {
	var todo model.Todo

	if err := r.db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *Repository) GetTodoByUser(userID string) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *Repository) GetTodoOfUserByStatus(userID string, done bool) ([]*model.Todo, error) {
	var todos []*model.Todo

	if err := r.db.Where("user_id = ? AND done = ?", userID, done).Find(&todos).Error; err != nil {
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

func (r *Repository) UpdateTodoDone(todo model.Todo) (*model.Todo, error) {

	t := model.Todo{
		ID:   todo.ID,
		Done: todo.Done,
	}

	if err := r.db.Model(&todo).Where("id = ?", t.ID).Update("done", t.Done).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *Repository) UpdateTodoTask(todo model.Todo) (*model.Todo, error) {

	t := model.Todo{
		ID:   todo.ID,
		Task: todo.Task,
	}

	if err := r.db.Model(&todo).Where("id = ?", t.ID).Update("task", t.Task).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
