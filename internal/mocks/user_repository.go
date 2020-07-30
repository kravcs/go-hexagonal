package mocks

import (
	u "hexa/internal/user"

	mock "github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: name
func (_m *UserRepository) FindByName(name string) (*u.User, error) {
	ret := _m.Called(name)

	var r0 *u.User
	if rf, ok := ret.Get(0).(func(string) *u.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*u.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
