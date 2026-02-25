package services

import (
	"errors"
	"mime/multipart"
	"to-do/dto"
	"to-do/models"
	"to-do/repository"
	"to-do/utils"
)

type CategoryService interface {
	Create(req *dto.CreateCategoryRequest, file multipart.File, fileHeader *multipart.FileHeader) (*dto.CategoryResponse, error)
	FindAll(filter dto.CategoryFilter) ([]dto.CategoryResponse, error)
	FindByID(id uint) (*dto.CategoryResponse, error)
	Update(id uint, req *dto.UpdateCategoryRequest, file multipart.File, fileHeader *multipart.FileHeader) (*dto.CategoryResponse, error)
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

// ==== CREATE ====
func (s *categoryService) Create(req *dto.CreateCategoryRequest, file multipart.File, fileHeader *multipart.FileHeader) (*dto.CategoryResponse, error) {

	// 1. Name unique ekanligini tekshirish
	existing, _ := s.repo.FindByName(req.Name)

	if existing != nil {
		return nil, errors.New("Bu nomli category allaqachon mavjud")
	}

	// 2. Image yuklash (majburiy emas)
	imageUrl := ""
	if file != nil {
		url, err := utils.UploadImage(file, fileHeader, "categories")
		if err != nil {
			return nil, err
		}
		imageUrl = url
	}

	// 3. Category yaratish

	category := &models.Category{
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
		Image:       imageUrl,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, errors.New("category yaratishda xatolik")
	}

	return toCategoryResponse(category), nil
}

// ==== FIND ALL ====
func (s *categoryService) FindAll(filter dto.CategoryFilter) ([]dto.CategoryResponse, error) {
	categories, err := s.repo.FindAll(filter)

	if err != nil {
		return nil, errors.New("categorylarni olishda xatolik")
	}

	var response []dto.CategoryResponse
	for _, category := range categories {
		response = append(response, *toCategoryResponse(&category))
	}
	return response, nil
}

// ==== FIND BY ID ====

func (s *categoryService) FindByID(id uint) (*dto.CategoryResponse, error) {
	category, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("category topilmadi")
	}

	return toCategoryResponse(category), nil
}

// ==== UPDATE ====
func (s *categoryService) Update(id uint, req *dto.UpdateCategoryRequest, file multipart.File, fileHeader *multipart.FileHeader) (*dto.CategoryResponse, error) {
	// 1. Category mavjudligini tekshirish
	category, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("category topilmadi")
	}

	// 2. Faqat yuborilgan fieldlarni yangilash
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.Color != "" {
		category.Color = req.Color
	}

	// 3. Yangi image yuklash
	if file != nil && fileHeader != nil {
		// Eski imageni Cloudinarydan o'chirish
		if category.Image != "" {
			publicID := utils.ExtractPublicID(category.Image)
			utils.DeleteImage(publicID)
		}

		url, err := utils.UploadImage(file, fileHeader, "categories")
		if err != nil {
			return nil, err
		}
		category.Image = url
	}

	if err := s.repo.Update(category); err != nil {
		return nil, errors.New("category yangilashda xatolik")
	}

	return toCategoryResponse(category), nil
}

// ==================== DELETE ====================

func (s *categoryService) Delete(id uint) error {
	// 1. Category mavjudligini tekshirish
	category, err := s.repo.FindById(id)
	if err != nil {
		return errors.New("category topilmadi")
	}

	// 2. Cloudinarydan image o'chirish
	if category.Image != "" {
		publicID := utils.ExtractPublicID(category.Image)
		utils.DeleteImage(publicID)
	}

	return s.repo.Delete(id)
}

// ==== CategoryResponse ====
func toCategoryResponse(category *models.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Color:       category.Color,
		Image:       category.Image,
	}
}
