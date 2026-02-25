package repository

import (
	"to-do/dto"
	"to-do/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindAll(filter dto.CategoryFilter) ([]models.Category, error)
	FindById(id uint) (*models.Category, error)
	FindByName(name string) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll(filter dto.CategoryFilter) ([]models.Category, error) {
	var categories []models.Category

	query := r.db.Where("deleted_at IS NULL")

	// Search — name yoki description bo'yicha
	//LIKE case-sensitive, ILIKE esa katta-kichik harfga e'tibor bermaydi.
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	// Sort
	allowedSortFields := map[string]bool{
		"name":       true,
		"created_at": true,
	}

	sortBy := "created_at"
	if !allowedSortFields[filter.SortBy] {
		sortBy = filter.SortBy
	}

	sortDir := "desc"
	if filter.SortDir == "asc" {
		sortDir = "asc"
	}

	query = query.Order(sortBy + " " + sortDir)

	err := query.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) FindById(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("name = ? AND deleted_at IS NULL", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Exec("UPDATE categories SET deleted_at = NOW() WHERE id = ?", id).Error
}
