package service

import (
	"context"

	"mini-shop/internal/product/entity"
)

type Repo interface {
	UpdateProductStock(ctx context.Context, product entity.Product) (bool, error)
	FindByID(ctx context.Context, id string) (entity.Product, error)
}

type Service struct {
	repo Repo
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(ctx context.Context, id string) (FindByIDOutput, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return FindByIDOutput{}, err
	}
	return FindByIDOutput{
		ID:       product.ID(),
		Name:     product.Name(),
		SellerID: product.SellerID(),
	}, nil
}

func (s *Service) DecreaseStock(ctx context.Context, input DecreaseStockInput) error {
	product, err := s.repo.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}
	err = product.DecreaseStock(input.Quantity, input.SellerID)
	if err != nil {
		return err
	}
	isSuccess, err := s.repo.UpdateProductStock(ctx, product)
	if err != nil {
		return err
	}
	if !isSuccess {
		return InsufficientStockErr
	}
	return nil
}
