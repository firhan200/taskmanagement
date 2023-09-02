package services

import (
	"errors"
	"log"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/repositories"
)

type UserService struct {
	ur *repositories.UserRepository
}

var (
	userService *UserService
)

func NewUserService(ur *repositories.UserRepository) *UserService {
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
	res, err := us.ur.GetByEmailAddressAndPassword(email, pass)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *UserService) Register(
	name string,
	email string,
	pass string,
) (*data.UserSecure, error) {
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
