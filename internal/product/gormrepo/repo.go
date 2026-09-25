package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/product/dto"
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

func (gr *GormRepo) SaveNew(ctx context.Context, product entity.Product) error {
	err := gr.db.WithContext(ctx).Create(&ProductGorm{
		ID:       product.ID(),
		SellerID: product.SellerID(),
		Name:     product.Name(),
		Stock:    0,
	}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create new product with gorm")
	}
	return nil
}

func (gr *GormRepo) FindMany(ctx context.Context, query dto.FindManyQuery) ([]entity.Product, error) {
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
		)
		if err != nil {
			continue
		}
		productEntities = append(productEntities, productEntity)
	}
	return productEntities, nil
}

func (gr *GormRepo) UpdateProductStock(ctx context.Context, product entity.Product) (bool, error) {
	result := gr.db.WithContext(ctx).
		Model(&ProductGorm{}).
		Where("id = ?", product.ID()).
		Where("stock >= ?", product.StockToDecr()).
		Update("stock", gorm.Expr("stock - ?", product.StockToDecr()))
	if result.Error != nil {
		return false, errors.Wrap(result.Error, "failed to decrease stock by id product with gorm")
	}
	return result.RowsAffected > 0, nil
}

func (gr *GormRepo) FindByID(ctx context.Context, id string) (entity.Product, error) {
	var productData ProductGorm
	err := gr.db.WithContext(ctx).
		Where("id = ?", id).
		Take(&productData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Product{}, ProductNotFoundErr
	}
	if err != nil {
		return entity.Product{}, errors.Wrap(err, "failed to find by id product with gorm")
	}
	return entity.NewProduct(
		productData.ID,
		productData.SellerID,
		productData.Name,
	)
}
