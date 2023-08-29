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
