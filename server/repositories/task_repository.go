package repositories

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/firhan200/taskmanagement/data"
	"gorm.io/gorm"
)

type ITaskRepository interface {
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

type TaskRepository struct {
	db ITaskRepository
}

var (
	taskRepository *TaskRepository
)

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	if taskRepository != nil {
		return taskRepository
	}
	log.Println("init new task repository")
	taskRepository = &TaskRepository{
		db: db,
	}

	return taskRepository
}

func (tr *TaskRepository) GetAll(
	uid uint,
	cursor interface{},
	limit int,
	orderBy string,
	sort string,
	keyword string,
) ([]data.Task, error) {
	var (
		tasks     []data.Task
		whereArgs string
	)

	if orderBy == "" {
		orderBy = "created_at"
	}
	if sort == "" {
		sort = "desc"
	}
	if limit == 0 {
		limit = 10
	}

	// start paginate if cursor is not empty
	if cursor == "" || cursor == 0 {
		whereArgs = getWhatToSort(orderBy) + " is not null AND ?='' AND title ILIKE ?"
	} else {
		whereArgs = filterCondition(orderBy, sort, keyword, true)
	}

	res := tr.db.Where(&data.Task{UserId: uid}).
		Order(fmt.Sprintf("%s %s", getWhatToSort(orderBy), sort)).
		Limit(limit).
		Find(&tasks, whereArgs, cursor, searchRule(keyword))

	if res.Error != nil {
		return nil, res.Error
	}

	return tasks, nil
}

func (tr *TaskRepository) GetTotalByUserId(
	uid uint,
	keyword string,
) (int, error) {
	var total int64
	res := tr.db.Find(&[]data.Task{}, "user_id = ? AND title ILIKE ? ", uid, searchRule(keyword)).
		Count(&total)
	if res.Error != nil {
		return 0, res.Error
	}

	return int(total), nil
}

func (tr *TaskRepository) GetNextCursor(
	uid uint,
	lastTask *data.Task,
	limit int,
	orderBy string,
	sort string,
	keyword string,
) interface{} {
	if lastTask == nil {
		return ""
	}

	var (
		nextTask       data.Task
		nextTaskCursor interface{}
	)

	nextTaskCursor = getTaskCursorColumnByOrder(lastTask, orderBy)

	tr.db.Order(fmt.Sprintf("%s %s", orderBy, sort)).
		Limit(limit).
		Find(&nextTask, filterCondition(orderBy, sort, keyword, false), nextTaskCursor, searchRule(keyword))
	if nextTask.ID == 0 {
		return ""
	}

	return getTaskCursorColumnByOrder(&nextTask, orderBy)
}

func getTaskCursorColumnByOrder(task *data.Task, orderBy string) interface{} {
	if orderBy == "due_date" {
		return task.DueDateUtcUnix
	}

	return task.CreatedAtUtcUnix
}

func getWhatToSort(orderBy string) string {
	if orderBy == "due_date" {
		return "due_date_utc_unix"
	}

	return "created_at_utc_unix"
}

func isCondEqual(cond string, isEqual bool) string {
	if isEqual {
		return fmt.Sprintf("%s=", cond)
	}

	return fmt.Sprint(cond)
}

func filterCondition(orderBy string, sort string, keyword string, isEqual bool) string {
	var whereArgs string
	if sort == "desc" {
		//result example: created_at_utc_unix <= ?
		whereArgs += fmt.Sprintf("%s %s ?", getWhatToSort(orderBy), isCondEqual("<", isEqual))
	} else {
		//result example: created_at_utc_unix >= ?
		whereArgs += fmt.Sprintf("%s %s ?", getWhatToSort(orderBy), isCondEqual(">", isEqual))
	}

	//result example: created_at_utc_unix <= ? AND title ILIKE ?
	whereArgs += " AND title ILIKE ?"

	//ex: created_at_utc_unix <= ? AND title ILIKE ?
	return whereArgs
}

func searchRule(keyword string) string {
	keywordValue := ""
	if keyword != "" {
		keywordValue = strings.ToLower(keyword)
	}
	return "%" + keywordValue + "%"
}

func (tr *TaskRepository) FindById(
	id uint,
) (*data.Task, error) {
	var (
		task *data.Task
	)

	//check if task exist
	res := tr.db.First(&task, id)
	if res.RowsAffected < 0 {
		return nil, errors.New("task not found")
	}

	return task, nil
}

/*=============== MUTATION =============*/
func (tm *TaskRepository) Insert(
	uid uint,
	title string,
	desc string,
	dueDate time.Time,
) (*data.Task, error) {
	createdTask := &data.Task{
		UserId:      uid,
		Title:       title,
		Description: desc,
		DueDate:     dueDate,
	}
	res := tm.db.Create(createdTask)

	return createdTask, res.Error
}

func (tm *TaskRepository) Update(
	uid uint,
	id uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	var (
		task *data.Task
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

func (tm *TaskRepository) Remove(
	uid uint,
	id uint,
) error {
	var (
		task *data.Task
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
