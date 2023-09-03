package repositories

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/firhan200/taskmanagement/data"
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

func TestTaskRepository_GetAll(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"id", "title", "description", "due_date"}).AddRow(1, "Testing", "Body", time.Now())
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	res, err := tr.GetAll(1, "", 0, "", "", "")
	assert.NoError(t, err)
	assert.Len(t, res, 1)
}

func TestTaskRepository_GetTotalByUserId(t *testing.T) {
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
	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	mock.ExpectQuery(`SELECT count`).WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(10))

	res, err := tr.GetTotalByUserId(1, "")
	assert.NoError(t, err)
	assert.Equal(t, 10, res)
}

func TestTaskRepository_GetNextCursor_NoLastTask(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	nextCursor := tr.GetNextCursor(1, nil, 10, "", "", "")
	assert.Empty(t, nextCursor)
}

func TestTaskRepository_GetNextCursor_WithLastTask(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id", "created_at_utc_unix"}).
		AddRow(1, time.Now().AddDate(0, 0, 1).UnixNano())
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	nextCursor := tr.GetNextCursor(1, &data.Task{
		Model: gorm.Model{
			ID: 1,
		},
		CreatedAtUtcUnix: time.Now().UnixNano(),
	}, 10, "created_at", "desc", "")
	assert.NotEmpty(t, nextCursor)
}

func TestTaskRepository_FindById_NotFound(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	_, err := tr.FindById(1)
	assert.Empty(t, err)
}

func TestTaskRepository_FindById_Found(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	task, err := tr.FindById(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, task)
	assert.Equal(t, uint(1), task.ID)
}

func TestTaskRepository_Insert_Error(t *testing.T) {
	m := new(MockTaskDB)
	tr := &TaskRepository{
		db: m,
	}

	m.On("Create", mock.Anything).Return(&gorm.DB{
		Error: errors.New("internal server error"),
	})

	_, err := tr.Insert(1, "title", "desc", time.Now())
	assert.Error(t, err)
}

func TestTaskRepository_Insert_Success(t *testing.T) {
	m := new(MockTaskDB)
	tr := &TaskRepository{
		db: m,
	}

	m.On("Create", mock.Anything).Return(&gorm.DB{})

	task, err := tr.Insert(1, "title", "desc", time.Now())
	log.Println(task)
	assert.NoError(t, err)
	assert.NotEmpty(t, task)
}

func TestTaskRepository_Update(t *testing.T) {
	m := new(MockTaskDB)
	tr := &TaskRepository{
		db: m,
	}

	m.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{
		RowsAffected: 0,
	})

	_, err := tr.Update(1, 2, "title", "desc", time.Now())
	assert.Error(t, err)
}

func TestTaskRepository_Update_Unauthorized(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id", "user_id"}).AddRow(5, 4)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	_, err := tr.Update(1, 2, "title", "desc", time.Now())
	assert.Error(t, err)
}

func TestTaskRepository_Update_Success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id", "user_id"}).AddRow(1, 2)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	mock.ExpectExec(`UPDATE`).WillReturnResult(sqlmock.NewResult(3, 1))

	_, err := tr.Update(2, 3, "title", "desc", time.Now())
	assert.NoError(t, err)
}

func TestTaskRepository_Delete_NotFound(t *testing.T) {
	m := new(MockTaskDB)
	tr := &TaskRepository{
		db: m,
	}

	m.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{
		RowsAffected: 0,
	})

	err := tr.Remove(1, 2)
	assert.Error(t, err)
}

func TestTaskRepository_Delete_Unauthorized(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id", "user_id"}).AddRow(5, 4)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	err := tr.Remove(1, 2)
	assert.Error(t, err)
}

func TestTaskRepository_Delete_Success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})
	tr := &TaskRepository{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"id", "user_id"}).AddRow(1, 2)
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	mock.ExpectExec(`UPDATE`).WillReturnResult(sqlmock.NewResult(1, 1))

	err := tr.Remove(2, 3)
	assert.NoError(t, err)
}
