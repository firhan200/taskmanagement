package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/services"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasksByUserId(
	uid uint,
	cursor interface{},
	limit int,
	orderBy string,
	sort string,
	search string,
) (*services.Tasks, error) {
	args := m.Called(uid, cursor, limit, orderBy, sort, search)
	return args.Get(0).(*services.Tasks), args.Error(0)
}
func (m *MockTaskService) GetByIdAuthorize(
	uid uint,
	id uint,
) (*data.Task, error) {
	args := m.Called(uid, id)
	return args.Get(0).(*data.Task), args.Error(0)
}
func (m *MockTaskService) Create(
	uid uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	args := m.Called(uid, title, description, dueDate)
	return args.Get(0).(*data.Task), args.Error(1)
}
func (m *MockTaskService) Update(
	uid uint,
	id uint,
	title string,
	description string,
	dueDate time.Time,
) (*data.Task, error) {
	args := m.Called(uid, id, title, description, dueDate)
	return args.Get(0).(*data.Task), args.Error(1)
}
func (m *MockTaskService) Delete(
	uid uint,
	id uint,
) error {
	args := m.Called(uid, id)
	return args.Error(0)
}

func createDummyJwt(uid int) string {
	token, err := utils.GenerateToken(1)
	if err != nil {
		log.Fatal("cannot create jwt token")
	}

	return token
}

func TestTaskController_CreateTask_Success(t *testing.T) {
	m := new(MockTaskService)
	handler := &TaskHandler{
		taskService: m,
	}

	//create mock
	m.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{
		Model: gorm.Model{
			ID: 1,
		},
	}, nil)

	app := fiber.New()
	app.Post("/tasks", handler.CreateTask())

	//add json body
	body, _ := json.Marshal(dto.CreateTaskDto{
		Title:       "Task Title Here",
		Description: "Lorem ipsum dolor sir amet",
		DueDate:     time.Now().AddDate(0, 0, 4),
	})
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	//create user auth
	token := createDummyJwt(1)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func BenchmarkCreateTask(b *testing.B) {
	m := new(MockTaskService)
	handler := &TaskHandler{
		taskService: m,
	}

	//create mock
	m.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&data.Task{
		Model: gorm.Model{
			ID: 1,
		},
	}, nil)

	app := fiber.New()
	app.Post("/tasks", handler.CreateTask())

	//add json body
	body, _ := json.Marshal(dto.CreateTaskDto{
		Title:       "Task Title Here",
		Description: "Lorem ipsum dolor sir amet",
		DueDate:     time.Now().AddDate(0, 0, 4),
	})
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	//create user auth
	token := createDummyJwt(1)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	for i := 0; i < b.N; i++ {
		app.Test(req)
	}
}
