package category

import (
	"context"

	"PayWatcher/domain"
	"PayWatcher/model"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

// Create implements model.CategoryRepository.
func (cr categoryRepo) Create(ctx context.Context, userID int, params model.UpdateOrCreateCategory) (model.Category, error) {
	var category model.Category

	category.UserID = uint(userID)
	category.Name = params.Name
	category.Priority = params.Priority
	category.Recurrent = params.Recurrent
	category.Notify = params.Notify

	if err := cr.db.Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// Delete implements model.CategoryRepository.
func (cr categoryRepo) Delete(ctx context.Context, ID, userID int) (model.Category, error) {
	var category model.Category

	if err := cr.db.First(&category, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return category, err
	}

	if err := cr.db.Delete(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// GetAll implements model.CategoryRepository.
func (cr categoryRepo) GetAll(ctx context.Context, userID int) ([]model.Category, error) {
	categories := []model.Category{}
	if err := cr.db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// GetByID implements model.CategoryRepository.
func (cr categoryRepo) GetByID(ctx context.Context, ID, userID int) (model.Category, error) {
	var category model.Category

	if err := cr.db.First(&category, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return category, err
	}

	return category, nil
}

// Update implements model.CategoryRepository.
func (cr categoryRepo) Update(ctx context.Context, ID, userID int, params model.UpdateOrCreateCategory) (model.Category, error) {
	var category model.Category

	if err := cr.db.First(&category, "id = ? AND user_id = ?", ID, userID).Error; err != nil {
		return category, err
	}

	category.Name = params.Name
	category.Priority = params.Priority
	category.Recurrent = params.Recurrent
	category.Notify = params.Notify

	if err := cr.db.Save(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func New(db *gorm.DB) domain.CategoryRepository {
	return categoryRepo{
		db: db,
	}
}
