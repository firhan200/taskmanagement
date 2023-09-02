package services

import (
	"errors"
	"log"
	"time"

	"github.com/firhan200/taskmanagement/data"
)

type ITaskRepository interface {
	GetAll(
		uid uint,
		cursor interface{},
		limit int,
		orderBy string,
		sort string,
		keyword string,
	) ([]data.Task, error)
	GetNextCursor(
		uid uint,
		lastTask *data.Task,
		limit int,
		orderBy string,
		sort string,
		keyword string,
	) interface{}
	GetTotalByUserId(
		uid uint,
		keyword string,
	) (int, error)
	FindById(
		id uint,
	) (*data.Task, error)
	Insert(
		uid uint,
		title string,
		desc string,
		dueDate time.Time,
	) (*data.Task, error)
	Update(
		uid uint,
		id uint,
		title string,
		description string,
		dueDate time.Time,
	) (*data.Task, error)
	Remove(
		uid uint,
		id uint,
	) error
}

type TaskService struct {
	tr ITaskRepository
}

type Tasks struct {
	Data       []data.Task
	Cursor     interface{}
	Limit      int
	OrderBy    string
	Sort       string
	Search     string
	Total      int
	NextCursor interface{}
}

var (
	taskService *TaskService
)

func NewTaskService(tr ITaskRepository) *TaskService {
	if taskService != nil {
		return taskService
	}

	log.Println("init new task service")
	taskService = &TaskService{
		tr: tr,
	}
	return taskService
}

func (ts *TaskService) GetTasksByUserId(
	uid uint,
	cursor interface{},
	limit int,
	orderBy string,
	sort string,
	search string,
) (*Tasks, error) {
	tasks := &Tasks{
		Data:    []data.Task{},
		Cursor:  cursor,
		Limit:   limit,
		OrderBy: orderBy,
		Sort:    sort,
		Search:  search,
	}

	res, err := ts.tr.GetAll(
		uid,
		cursor,
		limit,
		orderBy,
		sort,
		search,
	)
	if err != nil {
		return tasks, err
	}
	tasks.Data = res

	if len(res) < 1 {
		return tasks, nil
	}

	//set next cursor
	lastTask := res[len(res)-1]
	nc := ts.tr.GetNextCursor(
		uid,
		&lastTask,
		limit,
		orderBy,
		sort,
		search,
	)
	tasks.NextCursor = nc

	//get total
	total, err := ts.tr.GetTotalByUserId(uid, search)
	if err != nil {
		return tasks, err
	}
	tasks.Total = total

	return tasks, nil
}

func (ts *TaskService) GetByIdAuthorize(
	uid uint,
	id uint,
) (*data.Task, error) {
	res, err := ts.tr.FindById(id)
	if err != nil {
		return nil, err
	}

	//check if authorize
	if res.UserId != uid {
		return nil, errors.New("unauthorized actions")
	}

	return res, nil
}

func (ts *TaskService) Create(
	uid uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	res, err := ts.tr.Insert(
		uid,
		title,
		description,
		dueDate,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ts *TaskService) Update(
	uid uint,
	id uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	res, err := ts.tr.Update(
		uid,
		id,
		title,
		description,
		dueDate,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ts *TaskService) Delete(
	uid uint,
	id uint,
) error {
	err := ts.tr.Remove(
		uid,
		id,
	)

	return err
}
