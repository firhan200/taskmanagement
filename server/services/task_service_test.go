package services

import (
	"errors"
	"testing"
	"time"

	"github.com/firhan200/taskmanagement/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRespository struct {
	mock.Mock
}

func (m *MockTaskRespository) GetAll(
	uid uint,
	cursor interface{},
	limit int,
	orderBy string,
	sort string,
	keyword string,
) ([]data.Task, error) {
	args := m.Called(uid, cursor, limit, orderBy, sort, keyword)
	return args.Get(0).([]data.Task), args.Error(1)
}

func (m *MockTaskRespository) GetNextCursor(
	uid uint,
	lastTask *data.Task,
	limit int,
	orderBy string,
	sort string,
	keyword string,
) interface{} {
	args := m.Called(uid, lastTask, limit, orderBy, sort, keyword)
	return args.Get(0).(interface{})
}

func (m *MockTaskRespository) GetTotalByUserId(
	uid uint,
	keyword string,
) (int, error) {
	args := m.Called(uid, keyword)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockTaskRespository) FindById(
	id uint,
) (*data.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*data.Task), args.Error(1)
}

func (m *MockTaskRespository) Insert(
	uid uint,
	title string,
	desc string,
	dueDate time.Time,
) (*data.Task, error) {
	args := m.Called(uid, title, desc, dueDate)
	return args.Get(0).(*data.Task), args.Error(1)
}

func (m *MockTaskRespository) Update(
	uid uint,
	id uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	args := m.Called(uid, id, title, description, dueDate)
	return args.Get(0).(*data.Task), args.Error(1)
}

func (m *MockTaskRespository) Remove(
	uid uint,
	id uint,
) error {
	args := m.Called(uid, id)
	return args.Error(0)
}

func TestTaskService_NewTaskService(t *testing.T) {
	m := new(MockTaskRespository)
	ts := NewTaskService(m)

	assert.NotNil(t, ts)

	ts2 := NewTaskService(m)
	assert.NotNil(t, ts2)
}

func TestTaskService_GetTasksByUserId_GetAllError(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("GetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]data.Task{}, errors.New("failed to fetch data"))
	_, err := ts.GetTasksByUserId(1, "", 10, "created_at", "desc", "")
	assert.Error(t, err)
}

func TestTaskService_GetTasksByUserId_GetAllEmpty(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("GetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]data.Task{}, nil)
	_, err := ts.GetTasksByUserId(1, "", 10, "created_at", "desc", "")
	assert.Error(t, err)
}

func TestTaskService_GetTasksByUserId_GetAllNotEmpty(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("GetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]data.Task{
		{
			Title: "task 1",
		},
	}, nil)
	m.On("GetNextCursor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("")
	m.On("GetTotalByUserId", mock.Anything, mock.Anything).Return(0, errors.New("internal server error"))
	_, err := ts.GetTasksByUserId(1, "", 10, "created_at", "desc", "")
	assert.Error(t, err)
}

func TestTaskService_GetTasksByUserId_Complete(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("GetAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]data.Task{
		{
			Title: "task 1",
		},
	}, nil)
	m.On("GetNextCursor", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("")
	m.On("GetTotalByUserId", mock.Anything, mock.Anything).Return(1, nil)
	tasks, err := ts.GetTasksByUserId(1, "", 10, "created_at", "desc", "")
	assert.NoError(t, err)
	assert.Len(t, tasks.Data, 1)
}

func TestTaskService_GetByIdAuthorize_FindError(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("FindById", mock.Anything).Return(&data.Task{}, errors.New("task not found"))

	_, err := ts.GetByIdAuthorize(1, 2)
	assert.Error(t, err)
}

func TestTaskService_GetByIdAuthorize_FindSuccess(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("FindById", mock.Anything).Return(&data.Task{}, nil)

	_, err := ts.GetByIdAuthorize(1, 2)
	assert.Error(t, err)
}

func TestTaskService_GetByIdAuthorize_Unauthorized(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("FindById", mock.Anything).Return(&data.Task{
		UserId: 3,
	}, nil)

	_, err := ts.GetByIdAuthorize(1, 2)
	assert.Error(t, err)
}

func TestTaskService_GetByIdAuthorize_Complete(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("FindById", mock.Anything).Return(&data.Task{
		UserId: 1,
	}, nil)

	task, err := ts.GetByIdAuthorize(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), task.UserId)
}

func TestTaskService_Create_Failed(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Insert", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{}, errors.New("internal server error"))

	_, err := ts.Create(1, "title", "desc", time.Now())
	assert.Error(t, err)
}

func TestTaskService_Create_Success(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Insert", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{}, nil)

	_, err := ts.Create(1, "title", "desc", time.Now())
	assert.NoError(t, err)
}

func TestTaskService_Update_Failed(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{}, errors.New("internal server error"))

	_, err := ts.Update(1, 2, "title", "desc", time.Now())
	assert.Error(t, err)
}

func TestTaskService_Update_Success(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{}, nil)

	_, err := ts.Update(1, 2, "title", "desc", time.Now())
	assert.NoError(t, err)
}

func TestTaskService_Delete_Failed(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Remove", mock.Anything, mock.Anything).Return(errors.New("internal server error"))

	err := ts.Delete(1, 2)
	assert.Error(t, err)
}

func TestTaskService_Delete_Success(t *testing.T) {
	m := new(MockTaskRespository)
	ts := &TaskService{
		tr: m,
	}

	m.On("Remove", mock.Anything, mock.Anything).Return(nil)

	err := ts.Delete(1, 2)
	assert.NoError(t, err)
}
