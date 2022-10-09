package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUseCase interface {
	CreateUser(data dto.CreateUser) (error)
	GetAll() (model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{userRepository: userRepo}
}

func (u *userUseCase) CreateUser(data dto.CreateUser) (error){
	return repository.UserRepository.CreateUser(u.userRepository, data)
}

func (u *userUseCase) GetAll() (model.User, error) {
	return repository.UserRepository.GetAll(u.userRepository)
}