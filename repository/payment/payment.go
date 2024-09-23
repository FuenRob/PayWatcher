package payment

import (
	"PayWatcher/config"
	"PayWatcher/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

// Create Method
func (pr *PaymentRepo) Create(ctx context.Context, userID int, params model.UpdateOrCreatePayment) (model.Payment, error) {
	var payment model.Payment

	date, err := time.Parse(config.DateFormat, params.ChargeDate)

	if err != nil {
		return payment, err
	}

	payment.UserID = uint(userID)
	payment.Name = params.Name
	payment.CategoryID = params.CategoryID
	payment.NetAmount = params.NetAmount
	payment.GrossAmount = params.GrossAmount
	payment.Deductible = params.Deductible
	payment.ChargeDate = date
	payment.Recurrent = params.Recurrent
	payment.PaymentType = params.PaymentType
	payment.Paid = params.Paid

	if err := pr.db.Create(&payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

// Delete Method
func (pr *PaymentRepo) Delete(ctx context.Context, ID, userID int) (model.Payment, error) {
	var payment model.Payment

	if err := pr.db.First(&payment, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return payment, err
	}

	if err := pr.db.Delete(&payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

// GetAll Method
func (pr *PaymentRepo) GetAll(ctx context.Context, userID int) ([]model.Payment, error) {
	payments := []model.Payment{}
	if err := pr.db.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByCategoryID Method
func (pr *PaymentRepo) GetByCategoryID(ctx context.Context, categoryID, userID int) ([]model.Payment, error) {
	payments := []model.Payment{}
	if err := pr.db.Where("user_id = ? AND category_id = ?", userID, categoryID).Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// GetByID Method
func (pr *PaymentRepo) GetByID(ctx context.Context, ID, userID int) (model.Payment, error) {
	var payment model.Payment

	if err := pr.db.First(&payment, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

// Update Method
func (pr *PaymentRepo) Update(ctx context.Context, ID, userID int, params model.UpdateOrCreatePayment) (model.Payment, error) {
	var payment model.Payment

	date, err := time.Parse(config.DateFormat, params.ChargeDate)
	if err != nil {
		return payment, err
	}

	if err := pr.db.First(&payment, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return payment, err
	}

	payment.Name = params.Name
	payment.CategoryID = params.CategoryID
	payment.NetAmount = params.NetAmount
	payment.GrossAmount = params.GrossAmount
	payment.ChargeDate = date
	payment.Deductible = params.Deductible
	payment.Recurrent = params.Recurrent
	payment.PaymentType = params.PaymentType
	payment.Paid = params.Paid

	if err := pr.db.Save(&payment).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

func New(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}
