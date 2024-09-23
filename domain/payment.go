package domain

import (
	"PayWatcher/model"
	"context"
)

type PaymentRepository interface {
	GetAll(ctx context.Context, userID int) ([]model.Payment, error)
	GetByCategoryID(ctx context.Context, categoryID, userID int) ([]model.Payment, error)
	GetByID(ctx context.Context, ID, userID int) (model.Payment, error)
	Create(ctx context.Context, userID int, params model.UpdateOrCreatePayment) (model.Payment, error)
	Update(ctx context.Context, ID, userID int, params model.UpdateOrCreatePayment) (model.Payment, error)
	Delete(ctx context.Context, ID, userID int) (model.Payment, error)
}
