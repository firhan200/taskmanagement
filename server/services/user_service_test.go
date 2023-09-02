package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/firhan200/taskmanagement/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmailAddressAndPassword(emailAddress string, password string) (*data.User, error) {
	args := m.Called(emailAddress, password)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(emailAddress string) (*data.User, error) {
	args := m.Called(emailAddress)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserRepository) Insert(name string, emailAddress string, password string) (*data.UserSecure, error) {
	args := m.Called(name, emailAddress, password)
	return args.Get(0).(*data.UserSecure), args.Error(1)
}

func TestUserService_Login(t *testing.T) {
	m := new(MockUserRepository)
	us := NewUserService(m)

	var nilUser *data.User

	//check if login failed
	call := m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(nilUser, errors.New("user not found"))
	_, err := us.Login("not@user.com", "password")
	assert.Error(t, err, errors.New("user not found"))

	//test login success
	call.Unset()
	m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(&data.User{
		EmailAddress: "valid@email.com",
	}, nil)
	u, err := us.Login("valid@email.com", "password")
	assert.Nil(t, err)
	assert.Equal(t, "valid@email.com", u.EmailAddress)

	//test pass empty
	call.Unset()
	m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(nilUser, nil)
	_, passwordErr := us.Login("valid@email.com", "")
	assert.NotNil(t, passwordErr)

	//test email empty
	call.Unset()
	m.On("GetByEmailAddressAndPassword", mock.Anything, mock.Anything).Return(nilUser, nil)
	_, emailErr := us.Login("", "password")
	assert.NotNil(t, emailErr)
}

func TestUserService_Register(t *testing.T) {
	m := new(MockUserRepository)

	var user *data.User
	var userSecure *data.UserSecure

	m.On("FindByEmail", mock.Anything, mock.Anything, mock.Anything).Return(user, errors.New("user not found"))
	m.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(userSecure, nil)

	type fields struct {
		ur IUserRepository
	}
	type args struct {
		name  string
		email string
		pass  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *data.UserSecure
		wantErr bool
	}{
		{
			name: "register failed",
			fields: fields{
				ur: m,
			},
			args: args{
				name:  "",
				email: "",
				pass:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "register failed",
			fields: fields{
				ur: m,
			},
			args: args{
				name:  "name",
				email: "",
				pass:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "register failed",
			fields: fields{
				ur: m,
			},
			args: args{
				name:  "name",
				email: "email",
				pass:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "register success",
			fields: fields{
				ur: m,
			},
			args: args{
				name:  "name",
				email: "email",
				pass:  "pass",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				ur: tt.fields.ur,
			}
			got, err := us.Register(tt.args.name, tt.args.email, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
