package data

import (
	"fmt"

	"github.com/firhan200/taskmanagement/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName     string
	EmailAddress string
	Password     string
}

func (u *User) GetByEmailAddressAndPassword() error {
	db := GetConnection()

	//encrypt password
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	db.Where(&User{
		EmailAddress: u.EmailAddress,
		Password:     hashed,
	}).First(&u)

	if u.ID == 0 {
		return fmt.Errorf("User not found")
	}

	return nil
}

func (u *User) Save() (*User, error) {
	db := GetConnection()

	res := db.Model(u).Create(&u)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, fmt.Errorf("Failed to insert: %s", res.Error)
	}

	return u, nil
}

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
