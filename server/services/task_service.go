package services

import (
	"errors"
	"log"
	"time"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/repositories"
)

type TaskService struct {
	tr *repositories.TaskRepository
}

var (
	taskService *TaskService
)

func NewTaskService(tr *repositories.TaskRepository) *TaskService {
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
) (*data.Tasks, error) {
	tasks := &data.Tasks{
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

	//set next cursor
	if len(res) > 0 {
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
	}

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
