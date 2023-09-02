package data

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserId           uint
	Title            string
	Description      string
	DueDate          time.Time
	CreatedAtUtcUnix int64
	DueDateUtcUnix   int64
	Status           string `gorm:"<-:false;-:migration"`
	DueDateLocal     string `gorm:"<-:false;-:migration"`
}

/*=================== HOOKS ====================*/
func (t *Task) AfterFind(tx *gorm.DB) (err error) {
	//fic status
	dueDate := t.DueDate.UTC()
	now := time.Now()
	if !dueDate.IsZero() {
		if dueDate.Before(now) {
			t.Status = "Overdue"
		} else if dueDate.Before(now.AddDate(0, 0, 7)) && dueDate.After(now) {
			t.Status = "Due soon"
		} else if dueDate.After(now) {
			t.Status = "Not urgent"
		}
	}

	return
}

/*=================== HOOKS ====================*/
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.DueDateUtcUnix = t.DueDate.UTC().UnixNano()
	t.CreatedAtUtcUnix = time.Now().UTC().UnixNano()

	return
}
