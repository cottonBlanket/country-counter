package service

import (
	"github.com/cottonBlanket/country-counter/internal/models/domain"
	"github.com/cottonBlanket/country-counter/internal/models/dto"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(dto *dto.UserCreateDto) error {
	user := domain.User{Name: dto.Name}
	return s.db.Create(user).Error
}

func (s *UserService) UpdateUser(dto *dto.UserUpdateDto) error {
	var user domain.User
	if err := s.db.First(&user, dto.Id).Error; err != nil {
		return err
	}
	user.Name = dto.Name
	return s.db.Save(&user).Error
}

func (s *UserService) FindUsers() ([]domain.User, error) {
	var users []domain.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
