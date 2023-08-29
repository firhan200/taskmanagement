package data

import (
	"fmt"
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
	Data       []Task
	Cursor     interface{}
	Limit      int
	OrderBy    string
	Sort       string
	Search     string
	Total      int
	NextCursor interface{}
}

func IsCondEqual(cond string, isEqual bool) string {
	if isEqual {
		return fmt.Sprintf("%s=", cond)
	}

	return fmt.Sprintf("%s", cond)
}

func FilterCondition(orderBy string, sort string, keyword string, isEqual bool) string {
	var whereArgs string
	if sort == "desc" {
		whereArgs += fmt.Sprintf("%s %s ?", orderBy, IsCondEqual("<", isEqual))
	} else {
		whereArgs += fmt.Sprintf("%s %s ?", orderBy, IsCondEqual(">", isEqual))
	}

	//filter searching
	whereArgs += fmt.Sprintf(" AND title LIKE ?")

	return whereArgs
}

func SearchRule(keyword string) string {
	return "%" + keyword + "%"
}

func (ts *Tasks) GetByUserId(uid uint) {
	db := GetConnection()

	if ts.OrderBy == "" {
		ts.OrderBy = "created_at"
	}

	if ts.Sort == "" {
		ts.Sort = "desc"
	}

	if ts.Limit == 0 {
		ts.Limit = 10
	}

	//only need where cursor if sort by asc and cursor is 0
	var whereArgs string
	if ts.Cursor == "" {
		whereArgs = ts.OrderBy + " is not null AND ?='' AND title LIKE ?"
	} else {
		whereArgs = FilterCondition(ts.OrderBy, ts.Sort, ts.Search, true)
	}

	db.Where(&Task{UserId: uid}).
		Order(fmt.Sprintf("%s %s", ts.OrderBy, ts.Sort)).
		Limit(ts.Limit).
		Find(&ts.Data, whereArgs, ts.Cursor, SearchRule(ts.Search))

	//get total
	var total int64
	db.Find(&[]Task{}, "user_id = ? ", uid).
		Count(&total)
	ts.Total = int(total)

	//get last data
	if len(ts.Data) < 1 {
		return
	}
	lastTask := ts.Data[len(ts.Data)-1]
	if lastTask.ID == 0 {
		return
	}

	var (
		nextTask       Task
		nextTaskCursor interface{}
	)
	if ts.OrderBy == "created_at" {
		nextTaskCursor = lastTask.CreatedAt
	} else if ts.OrderBy == "due_date" {
		nextTaskCursor = lastTask.DueDate
	}

	db.Order(fmt.Sprintf("%s,id %s", ts.OrderBy, ts.Sort)).
		Limit(ts.Limit).
		Find(&nextTask, FilterCondition(ts.OrderBy, ts.Sort, ts.Search, false), nextTaskCursor, SearchRule(ts.Search))
	if nextTask.ID == 0 {
		return
	}

	if ts.OrderBy == "created_at" {
		ts.NextCursor = nextTask.CreatedAt
	} else if ts.OrderBy == "due_date" {
		ts.NextCursor = nextTask.DueDate
	}
}

func GetTask(id uint, userId uint) (*Task, error) {
	db := GetConnection()
	var task Task
	res := db.First(&task, id)

	if res.Error != nil {
		return &task, res.Error
	}

	if task.UserId != userId {
		return &task, fmt.Errorf("Unauthorized action")
	}

	return &task, res.Error
}

func (t *Task) Save() (*Task, error) {
	db := GetConnection()
	res := db.Create(t)

	return t, res.Error
}

func (t *Task) Update() (*Task, error) {
	db := GetConnection()

	var task *Task
	res := db.First(&task, t.ID)
	if res.RowsAffected < 0 {
		return nil, res.Error
	}

	//check if authorize
	if task.UserId != t.UserId {
		return nil, fmt.Errorf("Unauthorized to make changes")
	}

	fmt.Print(t)

	//update
	db.Save(&t)

	return t, res.Error
}

func (t *Task) Delete() error {
	db := GetConnection()

	var task *Task
	res := db.First(&task, t.ID)
	if res.RowsAffected < 0 {
		return res.Error
	}

	//check if authorize
	if task.UserId != t.UserId {
		return fmt.Errorf("Unauthorized to make changes")
	}

	//update
	db.Delete(&t)

	return res.Error
}
