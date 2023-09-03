package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockTaskDB struct {
	mock.Mock
}

func (m *MockTaskDB) Create(value interface{}) (tx *gorm.DB) {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Save(value interface{}) (tx *gorm.DB) {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockTaskDB) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Limit(limit int) (tx *gorm.DB) {
	args := m.Called(limit)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Order(value interface{}) (tx *gorm.DB) {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Count(count *int64) (tx *gorm.DB) {
	args := m.Called(count)
	return args.Get(0).(*gorm.DB)
}
func (m *MockTaskDB) Where(query interface{}, args ...interface{}) (tx *gorm.DB) {
	margs := m.Called(query, args)
	return margs.Get(0).(*gorm.DB)
}

func TestTaskRepository_NewTaskRepository(t *testing.T) {
	m := new(MockTaskDB)
	tr := NewTaskRepository(m)
	assert.NotNil(t, tr)

	tr2 := NewTaskRepository(m)
	assert.NotNil(t, tr2)
}

func TestTaskRepository_GetAll_NoCursor(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	// fmt.Println(mock)
	rows := sqlmock.NewRows([]string{"Titlea", "Description"}).AddRow("Testing", "Body")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	_, err := tr.GetAll(1, "", 0, "", "", "")
	assert.NoError(t, err)
}
