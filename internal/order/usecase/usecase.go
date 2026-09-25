package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"mini-shop/internal/order/entity"
	"mini-shop/internal/order/dto"
	productSvc "mini-shop/internal/product/service"
)

type Repo interface {
	SaveNew(ctx context.Context, order entity.Order) error
	FindById(ctx context.Context, id string) (entity.Order, error)
	FindAllBySellerId(ctx context.Context, query dto.FindAllBySellerIdQuery) ([]entity.Order, error)
	UpdateStatus(ctx context.Context, order entity.Order) (bool, error)
}

type ProductSvc interface {
	FindByID(ctx context.Context, id string) (productSvc.FindByIDOutput, error)
	DecreaseStock(ctx context.Context, input productSvc.DecreaseStockInput) error
}

type Transaction interface {
	Repo() Repo
	ProductSvc() ProductSvc

	Commit() error
	Rollback() error
}

type TxManager interface {
	New(ctx context.Context) (Transaction, error)
}

type Usecase struct {
	repo       Repo
	productSvc ProductSvc
	txManager  TxManager
}

func NewUsecase(
	repo Repo,
	productSvc ProductSvc,
	txManager TxManager,
) *Usecase {
	return &Usecase{
		repo:       repo,
		productSvc: productSvc,
		txManager:  txManager,
	}
}

func (u *Usecase) MakeOrder(ctx context.Context, input dto.MakeOrderInput) error {
	product, err := u.productSvc.FindByID(ctx, input.ProductID)
	if err != nil {
		return err
	}
	order, err := entity.NewOrder(
		uuid.NewString(),
		input.CustomerID,
		input.ProductID,
		product.SellerID,
		product.Name,
		input.Quantity,
		time.Now(),
		entity.PendingStatus,
	)
	if err != nil {
		return err
	}
	err = u.repo.SaveNew(ctx, order)
	return err
}

func (u *Usecase) OrderList(ctx context.Context, input dto.OrderListInput) ([]dto.OrderListItem, error) {
	orders, err := u.repo.FindAllBySellerId(ctx, dto.FindAllBySellerIdQuery{
		SellerID: input.SellerID,
		Limit:    input.Limit,
		Offset:   input.Offset,
	})
	if err != nil {
		return nil, err
	}
	res := make([]dto.OrderListItem, 0, len(orders))
	for _, order := range orders {
		res = append(res, dto.OrderListItem{
			OrderID:     order.ID(),
			ProductName: order.ProductName(),
			Quantity:    order.Quantity(),
		})
	}
	return res, nil
}

func (u *Usecase) CompleteOrder(ctx context.Context, input dto.CompleteOrderInput) error {
	order, err := u.repo.FindById(ctx, input.OrderID)
	if err != nil {
		return err
	}
	tx, err := u.txManager.New(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = order.Complete(input.SellerID)
	if err != nil {
		return err
	}
	isSuccess, err := tx.Repo().UpdateStatus(ctx, order)
	if err != nil {
		return err
	}
	if !isSuccess {
		return InconsistentStatusChangesErr
	}
	err = tx.ProductSvc().DecreaseStock(ctx, productSvc.DecreaseStockInput{
		ID:       order.ProductID(),
		Quantity: order.Quantity(),
		SellerID: input.SellerID,
	})
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func (u *Usecase) CancelOrder(ctx context.Context, input dto.CancelOrderInput) error {
	order, err := u.repo.FindById(ctx, input.OrderID)
	if err != nil {
		return err
	}
	err = order.Cancel(input.SellerID)
	if err != nil {
		return err
	}
	isSuccess, err := u.repo.UpdateStatus(ctx, order)
	if err != nil {
		return err
	}
	if !isSuccess {
		return InconsistentStatusChangesErr
	}
	return nil
}
