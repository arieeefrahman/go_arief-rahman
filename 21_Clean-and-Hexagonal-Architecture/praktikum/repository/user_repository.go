package repository

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(data dto.CreateUser) (error)
	GetAll() (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(data dto.CreateUser) (error){
	return r.db.Create(&data).Error
}

func (r *userRepository) GetAll() (model.User, error) {
	var user model.User

	if err := r.db.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}