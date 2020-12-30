package service_test

import (
	"errors"
	"fmt"
	"github.com/rickb777/date"
	"github.com/stretchr/testify/assert"
	"github.com/yet-another-todo-list-golang/model/dto"
	"github.com/yet-another-todo-list-golang/model/entity"
	"os"
	"testing"
	"time"

	mocks "github.com/yet-another-todo-list-golang/mocks/repository"
	"github.com/yet-another-todo-list-golang/service"
)

var todos []entity.Todo

func TestMain(m *testing.M) {
	for i := 1; i < 5; i++ {
		todos = append(todos, entity.Todo{
			Base: entity.Base{
				ID: fmt.Sprintf("35c54eae-ffaf-4085-a545-bb66d2abc22%d", i),
			},
			Name:        fmt.Sprintf("Todo%d", i),
			Description: fmt.Sprintf("Todo%d", i),
			Status:      "ACTIVE",
			UserID:      "",
			User:        entity.User{},
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		})
	}

	os.Exit(m.Run())
}

func TestTodoFindAll(t *testing.T) {
	todoRepository := new(mocks.TodoRepository)

	testCases := []struct {
		name   string
		entity []entity.Todo
		err    error
	}{
		{
			name:   "Simple read",
			entity: todos,
			err:    nil,
		},
		{
			name:   "Unknown error such as database error",
			entity: []entity.Todo{},
			err:    errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todoRepository.On("FindAll").Return(tc.err, tc.entity)
			todoService := service.NewTodoService(todoRepository)
			err, got := todoService.FindAll()
			if err != nil {
				assert.EqualError(t, tc.err, err.Error())
				return
			}

			assert.Equal(t, got, todos)
		})
	}
}

func TestTodoFindOne(t *testing.T) {
	todoRepository := new(mocks.TodoRepository)
	testCases := []struct {
		name     string
		input    string
		expected entity.Todo
		err      error
	}{
		{
			name:     "Get one",
			input:    todos[1].ID,
			expected: todos[1],
			err:      nil,
		},
		{
			name:     "Record not found",
			input:    "8c113ede-a8f8-473a-b4f8-0216902538c0",
			expected: entity.Todo{},
			err:      errors.New("record not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todoRepository.On("FindOne", tc.input).Return(tc.err, tc.expected)
			todoService := service.NewTodoService(todoRepository)

			err, got := todoService.FindOne(tc.input)
			if err != nil {
				assert.EqualError(t, tc.err, err.Error())
				return
			}
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestTodoCreate(t *testing.T) {
	todoRepository := new(mocks.TodoRepository)

	testCases := []struct {
		name     string
		input    dto.CreateTodoDto
		expected error
		user     entity.User
	}{
		{
			name: "Create todo",
			input: dto.CreateTodoDto{
				Name:        "Todo1",
				Description: "Todo1",
			},
			user: entity.User{
				Base:        entity.Base{ID: "6c624b8a-1378-4a5d-a10b-c9bce7863c03"},
				Email:       "test@test.com",
				Password:    "test",
				DateOfBirth: date.DateString{},
			},
			expected: nil,
		},
		{
			name:  "Database error",
			input: dto.CreateTodoDto{},
			user: entity.User{
				Base:        entity.Base{ID: "6c624b8a-1378-4a5d-a10b-c9bce7863c03"},
				Email:       "test@test.com",
				Password:    "test",
				DateOfBirth: date.DateString{},
			},
			expected: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todo := entity.Todo{
				Base:        entity.Base{},
				Name:        tc.input.Name,
				Description: tc.input.Description,
				Status:      "",
				UserID:      "",
				User:        entity.User{},
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			}
			todoRepository.On("Create", todo, tc.user).Return(tc.expected)
			todoService := service.NewTodoService(todoRepository)

			err := todoService.Create(tc.input, tc.user)
			if err != nil {
				assert.EqualError(t, tc.expected, err.Error())
				return
			}

			assert.Equal(t, tc.expected, nil)
		})
	}
}

func TestTodoUpdate(t *testing.T) {

}

func TestTodoDelete(t *testing.T) {

}
