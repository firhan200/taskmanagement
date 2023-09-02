package data

import (
	"gorm.io/gorm"
)

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
