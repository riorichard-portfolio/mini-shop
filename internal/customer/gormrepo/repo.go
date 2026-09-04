package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/customer/entity"
)

type GormRepo struct {
	db *gorm.DB
}

func NewRepo(
	db *gorm.DB,
) *GormRepo {
	return &GormRepo{
		db: db,
	}
}

func (gr *GormRepo) SaveNew(ctx context.Context, customer entity.Customer) error {
	err := gr.db.WithContext(ctx).Create(&CustomerGorm{
		ID:             customer.ID(),
		Email:          customer.Email(),
		HashedPassword: customer.HashedPassword(),
	}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create new customer with gorm")
	}
	return nil
}

func (gr *GormRepo) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := gr.db.WithContext(ctx).
		Model(&CustomerGorm{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, errors.Wrap(err, "failed to count to check email exists with gorm")
	}
	return count > 0, nil
}

func (gr *GormRepo) FindByEmail(ctx context.Context, email string) (entity.Customer, error) {
	var customerData CustomerGorm
	err := gr.db.WithContext(ctx).
		Model(&CustomerGorm{}).
		Where("email = ?", email).
		First(&customerData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Customer{}, CustomerNotFoundErr
	}
	if err != nil {
		return entity.Customer{}, errors.Wrap(err, "failed to find first for find by email with gorm")
	}
	return entity.NewCustomer(
		customerData.ID,
		customerData.Email,
		customerData.HashedPassword,
	), nil
}
