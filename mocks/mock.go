package mocks

import (
	"github.com/stretchr/testify/mock"
	conf "DemystData/config"
	todo "DemystData/todo"
)

// MockTodo is a mock implementation of TodoInterface
type MockTodo struct {
	mock.Mock
}

type MockLogger struct {
	mock.Mock
}

// FetchEvenTodos is a mock implementation of FetchEvenTodos
func (m *MockTodo) FetchEvenTodos(n int, logger *MockLogger, config *conf.AppConfig) ([]todo.Todo, error) {
	args := m.Called(n, logger, config)
	return args.Get(0).([]todo.Todo), args.Error(1)
}

// MockFetchEvenTodos is a utility function to set up a mock implementation of FetchEvenTodos
func MockFetchEvenTodos(config conf.AppConfig, logger *MockLogger) *MockTodo {
	mockTodo := new(MockTodo)
	mockTodo.On("FetchEvenTodos", mock.Anything, mock.Anything, config).Return([]todo.Todo{}, nil)
	return mockTodo
}
