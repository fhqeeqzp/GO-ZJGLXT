package services

import (
	"go-wails-admin/internal/database"
	"go-wails-admin/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserList(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := database.GetDB().Model(&models.User{})
	db.Count(&total)

	err := db.Preload("Role").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error

	return users, total, err
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.GetDB().Preload("Role").First(&user, id).Error
	return &user, err
}

func (s *UserService) CreateUser(user *models.User) error {
	return database.GetDB().Create(user).Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	return database.GetDB().Save(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
	return database.GetDB().Delete(&models.User{}, id).Error
}
