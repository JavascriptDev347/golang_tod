package services

import (
	"errors"
	"to-do/dto"
	"to-do/models"
	"to-do/repository"
	"to-do/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {

	return &userService{
		repo: repo,
	}
}

// ==== REGISTER ====
func (s *userService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {

	// check the email
	existEmail, _ := s.repo.FindByEmail(req.Email)
	if existEmail != nil {
		return nil, errors.New("bu email allaqachon ro'yxatdan o'tgan")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("parolni hash qilishda xatolik")
	}

	// create user
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, errors.New("foydalanuvchi yaratishda xatolik")
	}

	// generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return utils.BuildAuthResponse(token, user), nil
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	// find the user with email
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email yoki parol noto'g'ri")
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email yoki parol noto'g'ri")
	}

	// generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return utils.BuildAuthResponse(token, user), nil
}
