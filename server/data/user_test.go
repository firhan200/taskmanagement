package data

import (
	"fmt"
	"testing"

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

func (m *MockIUserManager) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	mArgs := m.Called(dest, conds)
	return mArgs.Get(0).(*gorm.DB)
}

func TestGetByEmailAddressAndPassword_Failed(y *testing.T) {
	m := new(MockIUserManager)
	um := NewUserManager(m)

	u := User{
		FullName: "Test",
	}

	m.On("Find", &u, mock.Anything).Return(&gorm.DB{})

	emailAddress := "test@gmail.com"
	password := "123456"

	ua, _ := um.GetByEmailAddressAndPassword(emailAddress, password)

	// if err != nil {
	// 	log.Error(err.Error())
	// }

	fmt.Println("User", ua)
}
