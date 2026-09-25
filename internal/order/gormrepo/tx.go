package gormrepo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"mini-shop/internal/order/usecase"
	productRepo "mini-shop/internal/product/gormrepo"
	productSvc "mini-shop/internal/product/service"
)

type TxProcess struct {
	tx         *gorm.DB
	repo       *GormRepo
	productSvc *productSvc.Service
	isDone     bool
}

func (tp *TxProcess) Repo() usecase.Repo {
	return tp.repo
}

func (tp *TxProcess) ProductSvc() usecase.ProductSvc {
	return tp.productSvc
}

func (tp *TxProcess) Commit() error {
	if !tp.isDone {
		tp.isDone = true // if commit failed transaction can't be used again (only once commit/rollback)
		err := tp.tx.Commit().Error
		if err != nil {
			return errors.Wrap(err, "failed to commit transaction process order with gorm")
		}
	}
	return nil
}

func (tp *TxProcess) Rollback() error {
	if !tp.isDone {
		tp.isDone = true // if rollback failed transaction can't be used again (only once commit/rollback)
		err := tp.tx.Rollback().Error
		if err != nil {
			return errors.Wrap(err, "failed to rollback transaction process order with gorm")
		}
	}
	return nil
}

type TxManager struct {
	db *gorm.DB
}

func NewTxManager(
	db *gorm.DB,
) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (tm *TxManager) New(ctx context.Context) (usecase.Transaction, error) {
	tx := tm.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "failed to begin transaction process order with gorm")
	}
	productRepo := productRepo.NewRepo(tx)
	productSvc := productSvc.NewService(productRepo)
	return &TxProcess{
		tx:         tx,
		repo:       NewRepo(tx),
		productSvc: productSvc,
	}, nil
}
