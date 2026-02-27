package repository

import (
	"to-do/dto"
	"to-do/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	FindAll(userID uint, filter dto.TodoFilter) ([]models.Todo, error)
	FindByID(id uint, userID uint) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id uint, userID uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoResository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll(userID uint, filter dto.TodoFilter) ([]models.Todo, error) {

	var todos []models.Todo

	// query for find todos
	query := r.db.Where("todos.deleted_at IS NULL AND todos.user_id = ?", userID)

	// Search
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("todos.title ILIKE ? OR todos.description ILIKE ?", search, search)
	}

	// Status filter
	if filter.Status != "" {
		query = query.Where("todos.status = ?", filter.Status)
	}

	// Priority filter
	if filter.Priority != "" {
		query = query.Where("todos.priority = ?", filter.Priority)
	}

	// Sort
	allowedSortFields := map[string]bool{
		"title":      true,
		"status":     true,
		"priority":   true,
		"due_date":   true,
		"created_at": true,
	}

	sortBy := "created_at"
	sortDir := "desc"

	if filter.SortBy != "" && allowedSortFields[filter.SortBy] {
		sortBy = filter.SortBy
	}
	if filter.SortDir == "asc" {
		sortDir = "asc"
	}

	query = query.Order("todos." + sortBy + " " + sortDir)

	// Category ni ham yuklash
	err := query.Preload("Category").Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepository) FindByID(id, userId uint) (*models.Todo, error) {

	var todo models.Todo

	err := r.db.
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", id, userId).
		Preload("Category").
		First(&todo).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil

}

func (r *todoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint, userID uint) error {
	return r.db.Exec(
		"UPDATE todos SET deleted_at = NOW() WHERE id = ? AND user_id = ?",
		id, userID,
	).Error
}
