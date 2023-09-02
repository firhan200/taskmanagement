package services

import (
	"errors"
	"testing"

	"github.com/firhan200/taskmanagement/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmailAddressAndPassword(emailAddress string, password string) (*data.User, error) {
	args := m.Called(emailAddress, password)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(emailAddress string) (*data.User, error) {
	args := m.Called(emailAddress)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserRepository) Insert(name string, emailAddress string, password string) (*data.UserSecure, error) {
	args := m.Called(name, emailAddress, password)
	return args.Get(0).(*data.UserSecure), args.Error(1)
}

func TestUserService_NewUserService_Init(t *testing.T) {
	m := new(MockUserRepository)
	us := NewUserService(m)
	assert.NotNil(t, us)

	us2 := NewUserService(m)
	assert.NotNil(t, us2)
}

func TestUserService_Login_Failed(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	//check if login failed
	m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(&data.User{}, errors.New("user not found"))
	_, err := us.Login("not@user.com", "password")
	assert.Error(t, err)
}

func TestUserService_Login_ValidateParams(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	//check if login failed
	_, err := us.Login("", "password")
	assert.Error(t, err)

	_, err2 := us.Login("email@email.com", "")
	assert.Error(t, err2)
}

func TestUserService_Login_Success(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	//check if login failed
	m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(&data.User{
		EmailAddress: "valid@email.com",
	}, nil)
	u, err := us.Login("valid@email.com", "password")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserService_Register_ValidateParams(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	_, err := us.Register("", "", "")
	assert.Error(t, err)

	_, err2 := us.Register("name", "", "")
	assert.Error(t, err2)

	_, err3 := us.Register("name", "email@email.com", "")
	assert.Error(t, err3)
}

func TestUserService_Register_Failed(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	m.On("FindByEmail", mock.Anything).Return(&data.User{}, nil)

	_, err := us.Register("name", "valid@email.com", "password")
	assert.Error(t, err)
}

func TestUserService_Register_Success(t *testing.T) {
	m := new(MockUserRepository)
	us := &UserService{
		ur: m,
	}

	m.On("FindByEmail", mock.Anything).Return(&data.User{}, errors.New("user not found"))
	m.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(&data.UserSecure{}, nil)

	_, err := us.Register("name", "valid@email.com", "password")
	assert.NoError(t, err)
}
