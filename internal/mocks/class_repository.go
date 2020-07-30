package mocks

import (
	c "hexa/internal/class"

	mock "github.com/stretchr/testify/mock"
)

type ClassRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: redirect
func (_m *ClassRepository) Create(studioClass *c.Class) error {
	ret := _m.Called(studioClass)

	var r0 error
	if rf, ok := ret.Get(0).(func(*c.Class) error); ok {
		r0 = rf(studioClass)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *ClassRepository) FindByName(name string) (*c.Class, error) {
	ret := _m.Called(name)

	var r0 *c.Class
	if rf, ok := ret.Get(0).(func(string) *c.Class); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*c.Class)
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

func (m *ClassRepository) FindAll() (c.ClassesList, error) {
	ret := m.Called()

	var r0 c.ClassesList
	if rf, ok := ret.Get(0).(func() c.ClassesList); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(c.ClassesList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
