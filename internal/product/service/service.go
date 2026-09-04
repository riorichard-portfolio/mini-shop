package service

import (
	"context"

	"mini-shop/internal/product/entity"
)

type Repo interface {
	DecreaseStockByID(ctx context.Context, id string, qty int) (bool, error)
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
		ID:   product.ID(),
		Name: product.Name(),
	}, nil
}

func (s *Service) DecreaseStock(ctx context.Context, id string, qty int) error {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	err = product.DecreaseStock(qty)
	if err != nil {
		return err
	}
	isSuccess, err := s.repo.DecreaseStockByID(ctx, id, qty)
	if err != nil {
		return err
	}
	if !isSuccess {
		return InsufficientStockErr
	}
	return nil
}
