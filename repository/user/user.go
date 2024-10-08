package user

import (
	"PayWatcher/model"
	"context"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur *UserRepo) Create(ctx context.Context, params model.UpdateOrCreateUser) (model.User, error) {
	var user model.User

	user.Name = params.Name
	user.Email = params.Email
	user.UserName = params.UserName
	user.Password = params.Password

	if err := ur.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil

}

func (ur *UserRepo) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}

	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepo) GetByID(ctx context.Context, ID int) (model.User, error) {
	var user model.User

	if err := ur.db.First(&user, ID).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) Update(ctx context.Context, ID int, params model.UpdateOrCreateUser) (model.User, error) {
	var user model.User

	if err := ur.db.First(&user, ID).Error; err != nil {
		return user, err
	}

	user.Name = params.Name
	user.Email = params.Email
	user.UserName = params.UserName
	user.Password = params.Password

	if err := ur.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepo) Delete(ctx context.Context, ID int) (model.User, error) {
	var user model.User

	if err := ur.db.First(&user, ID).Error; err != nil {
		return user, err
	}

	if err := ur.db.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}
