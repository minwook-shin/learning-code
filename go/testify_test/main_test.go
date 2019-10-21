package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, 10, 10, "equal")
	assert.NotEqual(t, 1, 10, "not equal")
}

type MockObj struct {
	mock.Mock
}

func (mock *MockObj) testFunc(number int) (bool, error) {
	args := mock.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestFunc(t *testing.T) {
	testObj := new(MockObj)
	testObj.On("testFunc", 1).Return(true, nil)
	testObj.AssertExpectations(t)
}

type TestSuite struct {
	suite.Suite
	Variable int
}

func (suite *TestSuite) SetupTest() {
	suite.Variable = 1
}

func (suite *TestSuite) TestExample() {
	assert.Equal(suite.T(), 1, suite.Variable)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
