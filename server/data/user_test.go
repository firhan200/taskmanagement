package data

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockIUserManager struct {
	mock.Mock
}

func (m *MockIUserManager) Create(value interface{}) (tx *gorm.DB) {
	mArgs := m.Called(value)
	return mArgs.Get(0).(*gorm.DB)
}

func (m *MockIUserManager) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	mArgs := m.Called(dest, conds)
	return mArgs.Get(0).(*gorm.DB)
}

func (m *MockIUserManager) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	mArgs := m.Called(dest, conds)
	return mArgs.Get(0).(*gorm.DB)
}

func (m *MockIUserManager) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	mArgs := m.Called(query, args)
	return mArgs.Get(0).(*gorm.DB)
}

func TestGetByEmailAddressAndPassword_Failed(y *testing.T) {
	m := new(MockIUserManager)
	um := &UserManager{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	emailAddress := "test@gmail.com"
	password := "123456"

	u, err := um.GetByEmailAddressAndPassword(emailAddress, password)

	if err != nil {
		log.Error(err.Error())
	}

	fmt.Println(u)
}
