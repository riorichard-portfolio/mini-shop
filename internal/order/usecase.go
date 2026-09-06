package order

import (
	"context"
	"time"

	"github.com/google/uuid"

	"mini-shop/internal/order/entity"
	productSvc "mini-shop/internal/product/service"
)

type Repo interface {
	SaveNew(ctx context.Context, order entity.Order) error
	FindById(ctx context.Context, id string) (entity.Order, error)
	FindBySellerId(ctx context.Context, sellerId string) ([]entity.Order, error)
	UpdateById(ctx context.Context, order entity.Order) error
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

func (u *Usecase) MakeOrder(ctx context.Context, input MakeOrderInput) error {
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
		"PENDING",
	)
	if err != nil {
		return err
	}
	err = u.repo.SaveNew(ctx, order)
	return err
}

func (u *Usecase) CompleteOrder(ctx context.Context, input CompleteOrderInput) error {
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
	err = tx.Repo().UpdateById(ctx, order)
	if err != nil {
		return err
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
