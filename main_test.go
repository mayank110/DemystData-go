package main

import (
	"bytes"
	"testing"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	conf "DemystData/config"
	mocktodo "DemystData/mocks"
)

// mock implementation of logrus.Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	args := m.Called(fields)
	return args.Get(0).(*logrus.Entry)
}

func (m *MockLogger) Info(args ...interface{}) {
	m.Called(args...)
}

func (m *MockLogger) Fatal(args ...interface{}) {
	m.Called(args...)
}

type MainTestSuite struct {
	suite.Suite
	mockLogger *mocktodo.MockLogger
}

func (suite *MainTestSuite) SetupTest() {
	suite.mockLogger = new(mocktodo.MockLogger)
}

func (suite *MainTestSuite) TestMain() {
	// Redirect stdout for testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

	// Mock the FetchEvenTodos function
	mockConfig := conf.AppConfig{DefaultTodo: 5}
	suite.mockLogger.On("Info", 5).Once() // Expect logging of the number of Todos
	suite.mockLogger.On("WithFields", mock.Anything).Return(suite.mockLogger).Once()
	suite.mockLogger.On("Info", "TODO Information").Times(5) // Expect logging inside the loop

	mocktodo.MockFetchEvenTodos(mockConfig, suite.mockLogger)

	main()

	suite.mockLogger.AssertExpectations(suite.T())

	// Capture stdout for validation
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	out := <-outC

	// Validate stdout
	assert.Contains(suite.T(), out, "TODO Information", "Expected output containing TODO Information")

	// Reset mock for the next test
	suite.mockLogger.ExpectedCalls = nil

	// Test with invalid input
	suite.mockLogger.On("Fatal", mock.Anything).Once() // Expect a fatal error for invalid input

	main()

	suite.mockLogger.AssertExpectations(suite.T())
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
