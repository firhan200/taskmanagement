package data

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserId      int
	Title       string
	Description string
}
