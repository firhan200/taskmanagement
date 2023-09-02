package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserDB struct {
	mock.Mock
}

func (m *MockUserDB) Create(value interface{}) (tx *gorm.DB) {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockUserDB) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func TestUserRepository_NewUserRepository(t *testing.T) {
	m := new(MockUserDB)
	ur := NewUserRepository(m)
	assert.NotNil(t, ur)

	ur2 := NewUserRepository(m)
	assert.NotNil(t, ur2)
}

func TestUserRepository_GetByEmailAddressAndPassword_Failed(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		Error: errors.New("internal server error"),
	})

	_, err := ur.GetByEmailAddressAndPassword("email", "pass")
	assert.Error(t, err)
}

func TestUserRepository_GetByEmailAddressAndPassword_NotFound(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		Error:        nil,
		RowsAffected: 0,
	})

	_, err := ur.GetByEmailAddressAndPassword("email", "pass")
	assert.Error(t, err)
}

func TestUserRepository_GetByEmailAddressAndPassword_Success(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		Error:        nil,
		RowsAffected: 1,
	})

	_, err := ur.GetByEmailAddressAndPassword("email", "pass")
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail_NotFound(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		RowsAffected: 0,
	})

	_, err := ur.FindByEmail("email")
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail_Found(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{
		RowsAffected: 1,
	})

	_, err := ur.FindByEmail("email")
	assert.NoError(t, err)
}

func TestUserRepository_Insert_Failed(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Create", mock.Anything).Return(&gorm.DB{
		Error: errors.New("internal server error"),
	})

	_, err := ur.Insert("name", "valid@email.com", "password")
	assert.Error(t, err)
}

func TestUserRepository_Insert_Failed_NoAffectedRows(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Create", mock.Anything).Return(&gorm.DB{
		Error:        nil,
		RowsAffected: 0,
	})

	_, err := ur.Insert("name", "valid@email.com", "password")
	assert.Error(t, err)
}

func TestUserRepository_Insert_Failed_Success(t *testing.T) {
	m := new(MockUserDB)
	ur := &UserRepository{
		db: m,
	}

	m.On("Create", mock.Anything).Return(&gorm.DB{
		Error:        nil,
		RowsAffected: 1,
	})

	_, err := ur.Insert("name", "valid@email.com", "password")
	assert.NoError(t, err)
}
