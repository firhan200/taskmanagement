package data

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockIUserManager struct {
	mock.Mock
}

func (m *MockIUserManager) Create(value interface{}) (tx *gorm.DB) {
	mArgs := m.Called(value)
	return mArgs.Get(0).(*gorm.DB)
}

func (m *MockIUserManager) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	mArgs := m.Called(dest, conds)
	return mArgs.Get(0).(*gorm.DB)
}

func TestUserManager_GetByEmailAddressAndPassword(t *testing.T) {
	m := new(MockIUserManager)
	m.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	type fields struct {
		db IUserManager
	}
	type args struct {
		emailAddress string
		password     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "test login success",
			fields: fields{
				db: m,
			},
			args: args{
				emailAddress: "test",
				password:     "123",
			},
			want: &User{
				EmailAddress: "test",
				Password:     "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			um := &UserManager{
				db: tt.fields.db,
			}
			got, err := um.GetByEmailAddressAndPassword(tt.args.emailAddress, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserManager.GetByEmailAddressAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserManager.GetByEmailAddressAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
