package domain

import (
	"PayWatcher/model"
	"context"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, ID int) (model.User, error)
	Create(ctx context.Context, params model.UpdateOrCreateUser) (model.User, error)
	Update(ctx context.Context, ID int, params model.UpdateOrCreateUser) (model.User, error)
	Delete(ctx context.Context, ID int) (model.User, error)
}
