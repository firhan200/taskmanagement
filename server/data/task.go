package data

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

type ITaskManager interface {
	Create(value interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)
	Order(value interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
}

type TaskManager struct {
	db ITaskManager
}

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

type Tasks struct {
	tm         *TaskManager
	Data       []Task
	Cursor     interface{}
	Limit      int
	OrderBy    string
	Sort       string
	Search     string
	Total      int
	NextCursor interface{}
}

func NewTaskManager(db *gorm.DB) *TaskManager {
	return &TaskManager{
		db: db,
	}
}

func (tm *TaskManager) GetTasks(
	uid uint,
	cursor interface{},
	limit int,
	orderBy string,
	sort string,
	search string,
) *Tasks {
	tasks := &Tasks{
		tm:      tm,
		Cursor:  cursor,
		Limit:   limit,
		OrderBy: orderBy,
		Sort:    sort,
		Search:  search,
	}

	fmt.Printf("ID USERNYA: %d", uid)

	//validate parameters
	tasks.ValidateParams()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		tasks.QueryPagination(uid)
		tasks.GetNextCursor()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tasks.CountTotalData(uid)
	}()

	wg.Wait()

	return tasks
}

// set default value
func (ts *Tasks) ValidateParams() {
	if ts.OrderBy == "" {
		ts.OrderBy = "created_at"
	}

	if ts.Sort == "" {
		ts.Sort = "desc"
	}

	if ts.Limit == 0 {
		ts.Limit = 10
	}
}

func getWhatToSort(orderBy string) string {
	if orderBy == "due_date" {
		return "due_date_utc_unix"
	}

	return "created_at_utc_unix"
}

func (ts *Tasks) QueryPagination(uid uint) {
	// only need where cursor if sort by asc and cursor is 0
	var whereArgs string
	if ts.Cursor == "" || ts.Cursor == 0 {
		whereArgs = getWhatToSort(ts.OrderBy) + " is not null AND ?='' AND title ILIKE ?"
	} else {
		whereArgs = FilterCondition(ts.OrderBy, ts.Sort, ts.Search, true)
	}

	ts.tm.db.Where(&Task{UserId: uid}).
		Order(fmt.Sprintf("%s %s", getWhatToSort(ts.OrderBy), ts.Sort)).
		Limit(ts.Limit).
		Find(&ts.Data, whereArgs, ts.Cursor, SearchRule(ts.Search))
}

func (ts *Tasks) CountTotalData(uid uint) {
	// get total
	var total int64
	ts.tm.db.Find(&[]Task{}, "user_id = ? AND title ILIKE ? ", uid, SearchRule(ts.Search)).
		Count(&total)
	ts.Total = int(total)
}

// only call this function after get main query data
func (ts *Tasks) GetNextCursor() {
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
		nextTaskCursor = lastTask.CreatedAtUtcUnix
	} else if ts.OrderBy == "due_date" {
		nextTaskCursor = lastTask.DueDateUtcUnix
	}

	ts.tm.db.Order(fmt.Sprintf("%s %s", ts.OrderBy, ts.Sort)).
		Limit(ts.Limit).
		Find(&nextTask, FilterCondition(ts.OrderBy, ts.Sort, ts.Search, false), nextTaskCursor, SearchRule(ts.Search))
	if nextTask.ID == 0 {
		return
	}

	if ts.OrderBy == "created_at" {
		ts.NextCursor = nextTask.CreatedAtUtcUnix
	} else if ts.OrderBy == "due_date" {
		ts.NextCursor = nextTask.DueDateUtcUnix
	}
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
		whereArgs += fmt.Sprintf("%s %s ?", getWhatToSort(orderBy), IsCondEqual("<", isEqual))
	} else {
		whereArgs += fmt.Sprintf("%s %s ?", getWhatToSort(orderBy), IsCondEqual(">", isEqual))
	}

	//filter searching
	whereArgs += fmt.Sprintf(" AND title ILIKE ?")

	return whereArgs
}

func SearchRule(keyword string) string {
	keywordValue := ""
	if keyword != "" {
		keywordValue = strings.ToLower(keyword)
	}
	return "%" + keywordValue + "%"
}

func (tm *TaskManager) GetSingleTask(uid uint, id uint) (*Task, error) {
	var (
		task *Task
	)

	//check if task exist
	res := tm.db.First(&task, id)
	if res.RowsAffected < 0 {
		return nil, errors.New("task not found")
	}

	//check if authorize
	if task.UserId != uid {
		return nil, fmt.Errorf("unauthorized to make changes")
	}

	return task, nil
}

/*=============== MUTATION =============*/
func (tm *TaskManager) Save(
	uid uint,
	title string,
	desc string,
	dueDate time.Time,
) (*Task, error) {
	createdTask := &Task{
		UserId:      uid,
		Title:       title,
		Description: desc,
		DueDate:     dueDate,
	}
	res := tm.db.Create(createdTask)

	return createdTask, res.Error
}

func (tm *TaskManager) Update(
	uid uint,
	id uint,
	title string,
	description string,
	dueDate time.Time,
) (*Task, error) {
	var (
		task *Task
	)

	//check if task exist
	res := tm.db.First(&task, id)
	if res.RowsAffected < 0 {
		return nil, errors.New("task not found")
	}

	//check if error
	if res.Error != nil {
		return nil, res.Error
	}

	//check if authorize
	if task.UserId != uid {
		return nil, fmt.Errorf("unauthorized to make changes")
	}

	//update
	task.Title = title
	task.Description = description
	task.DueDate = dueDate

	tm.db.Save(task)

	return task, nil
}

func (tm *TaskManager) Remove(
	uid uint,
	id uint,
) error {
	var (
		task *Task
	)

	//check if task exist
	res := tm.db.First(&task, id)
	if res.RowsAffected < 0 {
		return errors.New("task not found")
	}

	if res.Error != nil {
		return res.Error
	}

	//check if authorize
	if task.UserId != uid {
		return errors.New("unauthorized to make changes")
	}

	//update
	tm.db.Delete(&task)

	return res.Error
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
