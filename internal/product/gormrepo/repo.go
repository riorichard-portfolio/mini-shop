package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/product"
	"mini-shop/internal/product/entity"
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

func (gr *GormRepo) SaveNew(ctx context.Context, product *entity.Product) error {
	err := gr.db.WithContext(ctx).Create(&ProductGorm{
		ID:       product.ID(),
		SellerID: product.SellerID(),
		Name:     product.Name(),
		Stock:    product.Stock(),
	}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create new product with gorm")
	}
	return nil
}

func (gr *GormRepo) FindMany(ctx context.Context, query *product.FindManyQuery) ([]entity.Product, error) {
	var productsData []ProductGorm
	err := gr.db.WithContext(ctx).
		Model(&ProductGorm{}).
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&productsData).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to find many products with gorm")
	}
	productEntities := make([]entity.Product, 0, len(productsData))
	for _, product := range productsData {
		productEntity, err := entity.NewProduct(
			product.ID,
			product.SellerID,
			product.Name,
			product.Stock,
		)
		if err != nil {
			continue
		}
		productEntities = append(productEntities, *productEntity)
	}
	return productEntities, nil
}
