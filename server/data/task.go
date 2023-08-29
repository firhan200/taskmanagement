package data

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserId      uint
	Title       string
	Description string
	DueDate     time.Time
}

type Tasks struct {
	Data   []Task
	Cursor int
	Total  int
}

func (ts *Tasks) GetByUserId(uid uint) {
	db := GetConnection()

	db.Model(&Task{}).Find(&ts.Data, "user_id = ? AND id > ?", uid, ts.Cursor)
}

func GetSingleTask(id uint, uid int) (Task, error) {
	db := GetConnection()
	var task Task
	res := db.First(&task, id)
	return task, res.Error
}

func (t *Task) Save() (uint, error) {
	db := GetConnection()
	res := db.Create(t)

	return t.ID, res.Error
}
