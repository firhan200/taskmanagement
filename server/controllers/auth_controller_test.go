package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Login(email string, pass string) (*data.User, error) {
	args := m.Called(email, pass)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserService) Register(name string, email string, pass string) (*data.UserSecure, error) {
	args := m.Called(name, email, pass)
	return args.Get(0).(*data.UserSecure), args.Error(1)
}

func TestLoginHandler_Login_Failed_EmptyBody(t *testing.T) {
	m := new(MockUserService)
	handler := &LoginHandler{
		userService: m,
	}

	app := fiber.New()
	app.Post("/login", handler.Login())
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Add("Content-Type", "application/json")

	res, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestLoginHandler_Login_Failed_UserNotFound(t *testing.T) {
	m := new(MockUserService)
	handler := &LoginHandler{
		userService: m,
	}

	//create mock
	m.On("Login", mock.Anything, mock.Anything).Return(&data.User{}, errors.New("user not found"))

	app := fiber.New()
	app.Post("/login", handler.Login())
	//add json body
	body, _ := json.Marshal(dto.LoginDto{
		EmailAddress: "invalid@email.com",
		Password:     "password",
	})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	res, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestLoginHandler_Login_Success(t *testing.T) {
	m := new(MockUserService)
	handler := &LoginHandler{
		userService: m,
	}

	//create mock
	m.On("Login", mock.Anything, mock.Anything).Return(&data.User{
		Model: gorm.Model{
			ID: 1,
		},
	}, nil)

	app := fiber.New()
	app.Post("/login", handler.Login())
	//add json body
	body, _ := json.Marshal(dto.LoginDto{
		EmailAddress: "invalid@email.com",
		Password:     "password",
	})
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	res, _ := app.Test(req)

	bodyRes, _ := io.ReadAll(res.Body)
	type Resp struct {
		Token string `json:"token"`
	}
	var jsonRes Resp
	encodeErr := json.Unmarshal(bodyRes, &jsonRes)
	if encodeErr != nil {
		log.Fatal("error encoding ")
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, jsonRes.Token)
}
