package services

import (
	"errors"
	"log"

	"github.com/firhan200/taskmanagement/data"
)

type IUserRepository interface {
	GetByEmailAddressAndPassword(emailAddress string, password string) (*data.User, error)
	FindByEmail(emailAddress string) (*data.User, error)
	Insert(name string, emailAddress string, password string) (*data.UserSecure, error)
}

type UserService struct {
	ur IUserRepository
}

var (
	userService *UserService
)

func NewUserService(ur IUserRepository) *UserService {
	if userService != nil {
		return userService
	}

	log.Println("init new user service")
	userService = &UserService{
		ur: ur,
	}
	return userService
}

func (us *UserService) Login(
	email string,
	pass string,
) (*data.User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if pass == "" {
		return nil, errors.New("password cannot be empty")
	}

	res, err := us.ur.GetByEmailAddressAndPassword(email, pass)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (us *UserService) Register(
	name string,
	email string,
	pass string,
) (*data.UserSecure, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if pass == "" {
		return nil, errors.New("pass cannot be empty")
	}

	_, err := us.ur.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already taken")
	}

	res, err := us.ur.Insert(
		name,
		email,
		pass,
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}
