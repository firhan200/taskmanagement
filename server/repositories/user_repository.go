package repositories

import (
	"errors"
	"log"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/utils"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(value interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}

type UserRepository struct {
	db IUserRepository
}

var (
	userRepository *UserRepository
)

func NewUserRepository(db IUserRepository) *UserRepository {
	if userRepository != nil {
		return userRepository
	}
	log.Println("init new user repository")
	userRepository = &UserRepository{
		db: db,
	}

	return userRepository
}

func (um *UserRepository) GetByEmailAddressAndPassword(
	emailAddress string,
	password string,
) (*data.User, error) {
	//encrypt password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	var u data.User

	res := um.db.Find(&u, &data.User{
		EmailAddress: emailAddress,
		Password:     hashed,
	})

	if res.Error != nil {
		return nil, res.Error
	}

	if u.EmailAddress == "" {
		return nil, errors.New("user not found")
	}

	return &u, nil
}

func (um *UserRepository) FindByEmail(
	emailAddress string,
) (*data.User, error) {
	var existedUser *data.User

	um.db.Find(&existedUser, &data.User{
		EmailAddress: emailAddress,
	})

	if existedUser.ID == 0 {
		return nil, errors.New("user not found")
	}

	return existedUser, nil
}

func (um *UserRepository) Insert(
	fullName string,
	emailAddress string,
	password string,
) (*data.UserSecure, error) {
	user := &data.User{
		FullName:     fullName,
		EmailAddress: emailAddress,
		Password:     password,
	}

	res := um.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, errors.New("failed to create user")
	}

	return &data.UserSecure{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress,
	}, nil
}
