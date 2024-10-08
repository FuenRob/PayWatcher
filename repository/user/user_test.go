package user

import (
	"PayWatcher/model"
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	// settings mock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	// open connection with gorm
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, mock
}

func TestGetAll(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to get all users
	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "Email", "UserName", "Password"},
	).AddRows(
		[][]driver.Value{
			{1, "User 1", "user@user.com", "UserName", "123456"},
			{2, "User 2", "user2@user.com", "UserName2", "123456"},
		}...,
	)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	ctx := context.Background()

	// new user repository
	repo := New(testDB)
	users, err := repo.GetAll(ctx)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected len(users) == 2, got %v\n", len(users))
	}
}

func TestGetByID(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to get user by id
	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "Email", "UserName", "Password"},
	).AddRow(
		1, "User 1", "user@user.com", "UserName", "123456",
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(rows)

	ctx := context.Background()
	// new user repository
	repo := New(testDB)
	userID := 1
	user, err := repo.GetByID(ctx, userID)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if user.ID != uint(userID) {
		t.Fatalf("expected user ID == %d, got %d\n", userID, user.ID)
	}
}

func TestNotFoundGetByID(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to not found get user by id
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

	// new user repository
	repo := New(testDB)

	ctx := context.Background()
	userID := 1

	user, err := repo.GetByID(ctx, userID)

	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected err == %v, got %v\n", sql.ErrNoRows, err)
	}

	if user.ID != 0 {
		t.Fatalf("expected user ID == 0, got %v\n", user.ID)
	}
}

func TestCreateUser(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to create category
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WithArgs("User 1", "user@user.com", "UserName", "123456").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// new user repository
	repo := New(testDB)

	ctx := context.Background()

	params := model.UpdateOrCreateUser{
		Name:     "User 1",
		Email:    "user@user.com",
		UserName: "UserName",
		Password: "123456",
	}

	user, err := repo.Create(ctx, params)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if user.ID != 1 {
		t.Fatalf("expected user ID == 1, got %v\n", user.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to update user
	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "Email", "UserName", "Password"},
	).AddRow(
		1, "User 1", "user@user.com", "UserName", "123456",
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WithArgs("User 2", "user2@user.com", "UserName2", "1234567", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// new user repository
	repo := New(testDB)
	ctx := context.Background()
	ID := 1
	params := model.UpdateOrCreateUser{
		Name:     "User 2",
		Email:    "user2@user.com",
		UserName: "UserName2",
		Password: "1234567",
	}

	user, err := repo.Update(ctx, ID, params)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if user.ID != 1 {
		t.Fatalf("expected user ID == 1, got %v\n", user.ID)
	}
}

func TestDeleteUser(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	// mock to delete user
	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "Email", "UserName", "Password"},
	).AddRow(
		1, "User 1", "user@user.com", "UserName", "123456",
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 1).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// new user repository
	repo := New(testDB)
	ctx := context.Background()
	ID := 1
	user, err := repo.Delete(ctx, ID)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if user.ID != 1 {
		t.Fatalf("expected user ID == 1, got %v\n", user.ID)
	}
}
