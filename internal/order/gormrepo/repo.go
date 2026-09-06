package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/order"
	"mini-shop/internal/order/entity"
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

func (gr *GormRepo) SaveNew(ctx context.Context, order entity.Order) error {
	err := gr.db.WithContext(ctx).Create(&OrderGorm{
		ID:          order.ID(),
		CustomerID:  order.CustomerID(),
		ProductID:   order.ProductID(),
		SellerID:    order.SellerID(),
		ProductName: order.ProductName(),
		Quantity:    order.Quantity(),
		Status:      order.CurrentStatus(),
		CreatedAt:   order.CreatedAt(),
	}).Error
	if err != nil {
		return errors.Wrap(err, "failed to create new order with gorm")
	}
	return nil
}

func (gr *GormRepo) FindById(ctx context.Context, id string) (entity.Order, error) {
	var orderData OrderGorm
	err := gr.db.WithContext(ctx).
		Model(&OrderGorm{}).
		Where("id = ?", id).
		Take(&orderData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Order{}, OrderNotFoundErr
	}
	if err != nil {
		return entity.Order{}, errors.Wrap(err, "failed find by id order with gorm")
	}
	return entity.NewOrder(
		orderData.ID,
		orderData.CustomerID,
		orderData.ProductID,
		orderData.SellerID,
		orderData.ProductName,
		orderData.Quantity,
		orderData.CreatedAt,
		orderData.Status,
	)
}

func (gr *GormRepo) FindAllBySellerId(ctx context.Context, query order.FindAllBySellerIdQuery) ([]entity.Order, error) {
	var ordersData []OrderGorm
	err := gr.db.WithContext(ctx).
		Where("seller_id = ?", query.SellerID).
		Limit(query.Limit).
		Offset(query.Offset).
		Find(&ordersData).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed find all by seller id order with gorm")
	}
	orderEntities := make([]entity.Order, 0, len(ordersData))
	for _, order := range ordersData {
		orderEntity, err := entity.NewOrder(
			order.ID,
			order.CustomerID,
			order.ProductID,
			order.SellerID,
			order.ProductName,
			order.Quantity,
			order.CreatedAt,
			order.Status,
		)
		if err != nil {
			continue
		}
		orderEntities = append(orderEntities, orderEntity)
	}
	return orderEntities, nil
}

func (gr *GormRepo) UpdateStatus(ctx context.Context, order entity.Order) (bool, error) {
	statusTo, err := order.StatusTo()
	if err != nil {
		return false, err
	}
	result := gr.db.WithContext(ctx).
		Model(&OrderGorm{}).
		Where("id = ?", order.ID()).
		Where("status = ?", order.CurrentStatus()).
		Update("status", statusTo)
	if result.Error != nil {
		return false, errors.Wrap(result.Error, "failed update status order with gorm")
	}
	return result.RowsAffected > 0, nil
}
