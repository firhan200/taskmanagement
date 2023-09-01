package data

import (
	"errors"
	"fmt"

	"github.com/firhan200/taskmanagement/utils"
	"gorm.io/gorm"
)

type IUserManager interface {
	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
}

type UserManager struct {
	db IUserManager
}

type User struct {
	gorm.Model
	FullName     string
	EmailAddress string
	Password     string
}

type UserSecure struct {
	ID           uint
	FullName     string
	EmailAddress string
	Password     string
}

func NewUserManager(db IUserManager) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (um *UserManager) GetByEmailAddressAndPassword(
	emailAddress string,
	password string,
) (*User, error) {
	var u User

	//encrypt password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	um.db.Where(&User{
		EmailAddress: emailAddress,
		Password:     hashed,
	}).First(&u)

	if u.ID == 0 {
		return nil, errors.New("user not found")
	}

	return &u, nil
}

func (um *UserManager) Register(
	fullName string,
	emailAddress string,
	password string,
) (*UserSecure, error) {
	user := &User{
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

	return &UserSecure{
		ID:           user.ID,
		FullName:     user.FullName,
		EmailAddress: user.EmailAddress,
	}, nil
}

/*======================== HOOKS ============================*/
func (u *User) BeforeSave(db *gorm.DB) error {
	//check if email already taken
	var existedUser User
	db.Where((&User{
		EmailAddress: u.EmailAddress,
	})).Find(&existedUser)
	if existedUser.ID != 0 {
		return fmt.Errorf("Email Already Taken")
	}

	//turn password into hash
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPass

	return nil
}
