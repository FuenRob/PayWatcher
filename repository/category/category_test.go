package category

import (
	"context"
	"database/sql"
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
