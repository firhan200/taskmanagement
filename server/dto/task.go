package dto

import "time"

type CreateTaskDto struct {
	Title       string
	Description string
	DueDate     time.Time
}

type UpdateTaskDto struct {
	CreateTaskDto
}
