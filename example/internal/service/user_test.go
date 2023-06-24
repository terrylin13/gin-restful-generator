package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/terrylin13/gin-restful-generator/example/internal/model"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) CreateUser(user *model.User) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) GetUser(user *model.User, id string) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) GetAllUsers() ([]*model.User, error) {
	args := mock.Called()
	return []*model.User{}, args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := &model.User{Name: "John"}

	//Setup expectations
	mockRepo.On("CreateUser").Return(nil)

	testService := UserService{repo: mockRepo}
	err := testService.CreateUser(user)

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockRepository)
	user := &model.User{Name: "John"}

	//Setup expectations
	mockRepo.On("GetUser").Return(nil)

	testService := UserService{repo: mockRepo}
	err := testService.GetUser(user, "1")

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockRepository)

	//Setup expectations
	mockRepo.On("GetAllUsers").Return(nil)
	testService := UserService{repo: mockRepo}
	users, err := testService.GetAllUsers()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.NoError(t, err)
	t.Log(users)
	// assert.Equal(t, "John", user.Name)
}
