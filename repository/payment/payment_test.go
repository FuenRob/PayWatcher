package payment

import (
	"PayWatcher/config"
	"PayWatcher/model"
	"context"
	"database/sql/driver"
	"log"
	"testing"
	"time"

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

func TestGetPaymentByID(t *testing.T) {
	testDB, mock := newMockDB()

	date, _ := time.Parse(config.DateFormat, "23-09-2024")

	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "CategoryID", "UserID", "NetAmount", "GrossAmount", "Deductible", "ChargeDate", "Recurrent", "PaymentType", "Paid"},
	).AddRow(
		1, "Payment 1", 1, 3, 100.0, 120.0, 20.0, date, true, "Credit Card", true,
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 3, 1).WillReturnRows(rows)

	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 3

	payment, err := repo.GetByID(ctx, ID, userID)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if payment.ID != uint(ID) {
		t.Fatalf("expected payment ID == %d, got %d\n", ID, payment.ID)
	}

	if payment.UserID != uint(userID) {
		t.Fatalf("expected userID == %d, got %d\n", userID, payment.UserID)
	}
}

func TestGetPaymentByIDError(t *testing.T) {
	testDB, mock := newMockDB()

	mock.ExpectQuery("SELECT").WillReturnError(gorm.ErrRecordNotFound)

	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 3

	_, err := repo.GetByID(ctx, ID, userID)

	if err == nil {
		t.Fatalf("expected err != nil, got %v\n", err)
	}
}

func TestGetPayments(t *testing.T) {
	testDB, mock := newMockDB()

	date, _ := time.Parse(config.DateFormat, "23-09-2024")

	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "CategoryID", "UserID", "NetAmount", "GrossAmount", "Deductible", "ChargeDate", "Recurrent", "PaymentType", "Paid"},
	).AddRows(
		[][]driver.Value{
			{1, "Payment 1", 1, 3, 100.0, 120.0, 20.0, date, true, "Credit Card", true},
			{2, "Payment 2", 2, 3, 200.0, 220.0, 20.0, date, true, "Credit Card", true},
			{3, "Payment 3", 3, 3, 300.0, 320.0, 20.0, date, true, "Credit Card", true},
		}...,
	)

	mock.ExpectQuery("SELECT").WithArgs(3).WillReturnRows(rows)

	repo := New(testDB)

	ctx := context.Background()
	userID := 3

	payments, err := repo.GetAll(ctx, userID)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}

	if len(payments) != 3 {
		t.Fatalf("expected payments length == 3, got %d\n", len(payments))
	}

	if payments[0].UserID != uint(userID) {
		t.Fatalf("expected userID == %d, got %d\n", userID, payments[0].UserID)
	}
}

func TestCreatePayment(t *testing.T) {
	config.DateFormat = "02-01-2006"
	testDB, mock := newMockDB()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `payments`").WillReturnResult(sqlmock.NewResult(1, 3))
	mock.ExpectCommit()

	repo := New(testDB)

	ctx := context.Background()
	userID := 3
	payment := model.UpdateOrCreatePayment{
		Name:        "Payment 1",
		CategoryID:  1,
		NetAmount:   100.0,
		GrossAmount: 120.0,
		Deductible:  20.0,
		ChargeDate:  "23-09-2024",
		Recurrent:   true,
		PaymentType: "Credit Card",
		Paid:        true,
	}

	_, err := repo.Create(ctx, userID, payment)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}
}

func TestUpdatePayment(t *testing.T) {
	config.DateFormat = "02-01-2006"
	testDB, mock := newMockDB()

	date, _ := time.Parse(config.DateFormat, "23-09-2024")

	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "CategoryID", "UserID", "NetAmount", "GrossAmount", "Deductible", "ChargeDate", "Recurrent", "PaymentType", "Paid"},
	).AddRow(
		1, "Payment 1", 1, 3, 100.0, 120.0, 20.0, date, true, "Credit Card", true,
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 3, 1).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `payments`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 3
	payment := model.UpdateOrCreatePayment{
		Name:        "Payment 1",
		CategoryID:  1,
		NetAmount:   100.0,
		GrossAmount: 120.0,
		Deductible:  20.0,
		ChargeDate:  "23-09-2024",
		Recurrent:   true,
		PaymentType: "Credit Card",
		Paid:        true,
	}

	_, err := repo.Update(ctx, ID, userID, payment)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}
}

func TestDeletePayment(t *testing.T) {
	testDB, mock := newMockDB()

	date, _ := time.Parse(config.DateFormat, "23-09-2024")

	rows := sqlmock.NewRows(
		[]string{"ID", "Name", "CategoryID", "UserID", "NetAmount", "GrossAmount", "Deductible", "ChargeDate", "Recurrent", "PaymentType", "Paid"},
	).AddRow(
		1, "Payment 1", 1, 3, 100.0, 120.0, 20.0, date, true, "Credit Card", true,
	)

	mock.ExpectQuery("SELECT").WithArgs(1, 3, 1).WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `payments`").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	repo := New(testDB)

	ctx := context.Background()
	ID := 1
	userID := 3

	_, err := repo.Delete(ctx, ID, userID)

	if err != nil {
		t.Fatalf("expected err == nil, got %v\n", err)
	}
}
