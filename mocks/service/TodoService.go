// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dto "github.com/yet-another-todo-list-golang/model/dto"
	entity "github.com/yet-another-todo-list-golang/model/entity"

	mock "github.com/stretchr/testify/mock"
)

// TodoService is an autogenerated mock type for the TodoService type
type TodoService struct {
	mock.Mock
}

// Create provides a mock function with given fields: createDto, user
func (_m *TodoService) Create(createDto dto.CreateTodoDto, user entity.User) error {
	ret := _m.Called(createDto, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.CreateTodoDto, entity.User) error); ok {
		r0 = rf(createDto, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *TodoService) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *TodoService) FindAll() (error, []entity.Todo) {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	var r1 []entity.Todo
	if rf, ok := ret.Get(1).(func() []entity.Todo); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entity.Todo)
		}
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *TodoService) FindOne(id string) (error, entity.Todo) {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	var r1 entity.Todo
	if rf, ok := ret.Get(1).(func(string) entity.Todo); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(entity.Todo)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updateDto
func (_m *TodoService) Update(id string, updateDto dto.UpdateTodoDto) error {
	ret := _m.Called(id, updateDto)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, dto.UpdateTodoDto) error); ok {
		r0 = rf(id, updateDto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
