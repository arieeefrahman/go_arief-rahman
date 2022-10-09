package controller_test

import (
	_userMock "belajar-go-echo/controller/mocks"
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.Repository
	userService usecase.UserUseCase

	userDomain model.User
	userDomainInput dto.CreateUser
)

func TestGetAllUsers(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		userRepository.On("GetAll").Return([]model.User{userDomain}).Once()

		result, _ := userService.GetAll()

		assert.Equal(t, 1, result)
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		userRepository.On("GetAll").Return([]model.User{}).Once()

		result, _ := userService.GetAll()

		assert.Equal(t, 0, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		userRepository.On("Create", &userDomainInput).Return(userDomainInput).Once()

		result := userService.CreateUser(userDomainInput)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		userRepository.On("Create", &model.User{}).Return(model.User{}).Once()

		result := userService.CreateUser(&model.User{})

		assert.NotNil(t, result)
	})
}