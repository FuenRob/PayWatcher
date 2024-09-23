package domain

import (
	"context"

	"PayWatcher/model"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, userID int) ([]model.Category, error)
	GetByID(ctx context.Context, ID, userID int) (model.Category, error)
	Create(ctx context.Context, userID int, params model.UpdateOrCreateCategory) (model.Category, error)
	Update(ctx context.Context, ID, userID, int, params model.UpdateOrCreateCategory) (model.Category, error)
	Delete(ctx context.Context, ID, userID int) (model.Category, error)
}
