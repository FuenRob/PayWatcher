package category

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

func TestGetCategoryByID(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()
	// mock to get category by id
	rows := sqlmock.NewRows(
		[]string{"ID", "UserID", "Name", "Priority", "Recurrent", "Notify"},
	).AddRow(
		1, 1, "Category 1", 1, true, false,
	)
	mock.ExpectQuery("SELECT").WithArgs(1, 1, 1).WillReturnRows(rows)

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 1

	category, err := repo.GetByID(ctx, ID, userID)
	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if category.ID != uint(ID) {
		t.Fatalf("expected category ID == %d, got %d\n", ID, category.ID)
	}

	if category.UserID != uint(userID) {
		t.Fatalf("expected userID == %d, got %d\n", userID, category.ID)
	}
}

func TestNotFoundCategoryByID(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()
	//mock to not found get category by id
	mock.ExpectQuery("SELECT").WithArgs(2, 1, 1).WillReturnError(sql.ErrNoRows)

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	ID := 2
	userID := 1

	category, err := repo.GetByID(ctx, ID, userID)
	if err == nil {
		t.Fatalf("expected err != nil, got %v\n", err)
	}

	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected err == %v, got %v\n", sql.ErrNoRows, err)
	}

	if category.ID != 0 {
		t.Fatalf("expected category ID == 0, got %v\n", category.ID)
	}
}

func TestGetAllCategory(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()
	// mock to get all categories
	rows := sqlmock.NewRows(
		[]string{"ID", "UserID", "Name", "Priority", "Recurrent", "Notify"},
	).AddRows(
		[][]driver.Value{
			{1, 1, "Category 1", 1, true, false},
			{2, 1, "Category 2", 2, false, true},
		}...,
	)
	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	userID := 1

	categories, err := repo.GetAll(ctx, userID)
	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if len(categories) != 2 {
		t.Fatalf("expected len(categories) == 2, got %d\n", len(categories))
	}
}

func TestNotFoundAllCategory(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()
	// mock to not found get all categories
	mock.ExpectQuery("SELECT").WithArgs(2).WillReturnError(sql.ErrNoRows)

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	userID := 2

	categories, err := repo.GetAll(ctx, userID)
	if err == nil {
		t.Fatalf("expected err != nil, got %v\n", err)
	}

	if !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected err == %v, got %v\n", sql.ErrNoRows, err)
	}

	if categories != nil {
		t.Fatalf("expected categories == nil, got %v\n", categories)
	}
}

func TestCreateCategory(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()
	// mock to create category
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WithArgs(1, "Category 1", 1, true, false).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	userID := 1
	params := model.UpdateOrCreateCategory{
		Name:      "Category 1",
		Priority:  1,
		Recurrent: true,
		Notify:    false,
	}

	category, err := repo.Create(ctx, userID, params)
	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if category.ID != 1 {
		t.Fatalf("expected category ID == 1, got %d\n", category.ID)
	}
}

func TestUpdateCategory(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	rows := sqlmock.NewRows(
		[]string{"ID", "UserID", "Name", "Priority", "Recurrent", "Notify"},
	).AddRow(
		1, 1, "Category 1", 1, true, false,
	)
	mock.ExpectQuery("SELECT").WithArgs(1, 1, 1).WillReturnRows(rows)

	// mock to update category
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WithArgs("Category 1", 1, true, false, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 1
	params := model.UpdateOrCreateCategory{
		Name:      "Category 1",
		Priority:  1,
		Recurrent: true,
		Notify:    false,
	}

	category, err := repo.Update(ctx, ID, userID, params)
	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if category.ID != 1 {
		t.Fatalf("expected category ID == 1, got %d\n", category.ID)
	}
}

func TestDeleteCategory(t *testing.T) {
	// get connection and mock
	testDB, mock := newMockDB()

	rows := sqlmock.NewRows(
		[]string{"ID", "UserID", "Name", "Priority", "Recurrent", "Notify"},
	).AddRow(
		1, 1, "Category 1", 1, true, false,
	)
	mock.ExpectQuery("SELECT").WithArgs(1, 1, 1).WillReturnRows(rows)

	// mock to delete category
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// new category repository
	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 1

	category, err := repo.Delete(ctx, ID, userID)
	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if category.ID != 1 {
		t.Fatalf("expected category ID == 1, got %d\n", category.ID)
	}
}
