package services

import (
	"errors"
	"time"
	"to-do/dto"
	"to-do/models"
	"to-do/repository"
)

type TodoService interface {
	Create(userID uint, req *dto.CreateTodoRequest) (*dto.TodoResponse, error)
	FindAll(userID uint, filter dto.TodoFilter) ([]dto.TodoResponse, error)
	FindByID(userID uint, id uint) (*dto.TodoResponse, error)
	Update(userID uint, id uint, req *dto.UpdateTodoRequest) (*dto.TodoResponse, error)
	Delete(userID uint, id uint) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) Create(userID uint, req *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
	dueDate, err := parseDate(req.DueDate)
	if err != nil {
		return nil, err
	}

	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    models.TodoPriority(req.Priority),
		DueDate:     dueDate,
		CategoryID:  req.CategoryID,
		UserID:      userID,
		Status:      models.StatusPending,
	}

	// Priority default
	if todo.Priority == "" {
		todo.Priority = models.PriorityMedium
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, errors.New("todo yaratishda xatolik")
	}

	return toTodoResponse(todo), nil
}

func (s *todoService) FindAll(userID uint, filter dto.TodoFilter) ([]dto.TodoResponse, error) {
	todos, err := s.repo.FindAll(userID, filter)
	if err != nil {
		return nil, errors.New("todolarni olishda xatolik")
	}

	var response []dto.TodoResponse
	for _, todo := range todos {
		response = append(response, *toTodoResponse(&todo))
	}

	return response, nil
}

func (s *todoService) FindByID(userID uint, id uint) (*dto.TodoResponse, error) {
	todo, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("todo topilmadi")
	}

	return toTodoResponse(todo), nil
}

func (s *todoService) Update(userID uint, id uint, req *dto.UpdateTodoRequest) (*dto.TodoResponse, error) {
	// 1. Todo mavjudligini tekshirish
	todo, err := s.repo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("todo topilmadi")
	}

	// 2. Faqat yuborilgan fieldlarni yangilash
	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.Description != "" {
		todo.Description = req.Description
	}
	if req.Status != "" {
		todo.Status = models.TodoStatus(req.Status)
	}
	if req.Priority != "" {
		todo.Priority = models.TodoPriority(req.Priority)
	}

	if req.CategoryID != nil {
		todo.CategoryID = req.CategoryID
	}

	dueDate, err := parseDate(req.DueDate)
	if err != nil {
		return nil, err
	}
	if req.DueDate != nil {
		todo.DueDate = dueDate
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, errors.New("todo yangilashda xatolik")
	}

	return toTodoResponse(todo), nil
}

func (s *todoService) Delete(userID uint, id uint) error {
	// Todo mavjudligini tekshirish
	_, err := s.repo.FindByID(id, userID)
	if err != nil {
		return errors.New("todo topilmadi")
	}

	return s.repo.Delete(id, userID)
}

func parseDate(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, nil
	}

	formats := []string{
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	for _, format := range formats {
		t, err := time.Parse(format, *dateStr)
		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("due_date formati noto'g'ri. Misol: 2027-01-02 yoki 2027-01-02T15:04:05Z")
}

func toTodoResponse(todo *models.Todo) *dto.TodoResponse {
	response := &dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		Priority:    string(todo.Priority),
		DueDate:     todo.DueDate,
		UserID:      todo.UserID,
		CategoryID:  todo.CategoryID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	// Category yuklangan bo'lsa
	if todo.Category != nil {
		response.Category = &dto.CategoryResponse{
			ID:          todo.Category.ID,
			Name:        todo.Category.Name,
			Description: todo.Category.Description,
			Color:       todo.Category.Color,
			Image:       todo.Category.Image,
		}
	}

	return response
}
