package mocks

import (
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) Create(userDomain *model.User) model.User {
	ret := _m.Called(userDomain)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(*model.User) model.User); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	return r0
}

func (_m *Usecase) GetAll() []model.User {
	ret := _m.Called()

	var r0 []model.User
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	return r0
}