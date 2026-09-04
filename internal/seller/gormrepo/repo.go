package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/seller/entity"
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

func (gr *GormRepo) SaveNew(ctx context.Context, seller entity.Seller) error {
	err := gr.db.WithContext(ctx).Create(&SellerGorm{
		ID:             seller.ID(),
		Email:          seller.Email(),
		HashedPassword: seller.HashedPassword(),
	}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create new seller with gorm")
	}
	return nil
}

func (gr *GormRepo) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := gr.db.WithContext(ctx).
		Model(&SellerGorm{}).
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		return false, errors.Wrap(err, "failed to count to check email exists with gorm")
	}
	return count > 0, nil
}

func (gr *GormRepo) FindByEmail(ctx context.Context, email string) (entity.Seller, error) {
	var sellerData SellerGorm
	err := gr.db.WithContext(ctx).
		Model(&SellerGorm{}).
		Where("email = ?", email).
		First(&sellerData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Seller{}, SellerNotFoundErr
	}
	if err != nil {
		return entity.Seller{}, errors.Wrap(err, "failed to find first for find by email with gorm")
	}
	return entity.NewSeller(
		sellerData.ID,
		sellerData.Email,
		sellerData.HashedPassword,
	), nil
}
