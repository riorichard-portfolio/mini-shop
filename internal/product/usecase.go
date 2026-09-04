package product

import (
	"context"

	"github.com/google/uuid"

	"mini-shop/internal/product/entity"
)

type Repo interface {
	SaveNew(ctx context.Context, product entity.Product) error
	FindMany(ctx context.Context, query FindManyQuery) ([]entity.Product, error)
}

type Usecase struct {
	repo Repo
}

func NewUsecase(
	repo Repo,
) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) AddNewProduct(ctx context.Context, input AddNewProductInput) error {
	newProduct, err := entity.NewProduct(
		uuid.NewString(),
		input.SellerID,
		input.Name,
		input.Stock,
	)
	if err != nil {
		return err
	}
	err = u.repo.SaveNew(ctx, newProduct)
	return err
}

func (u *Usecase) BrowseProducts(ctx context.Context, input BrowseProductsInput) ([]ProductItem, error) {
	productEntities, err := u.repo.FindMany(ctx, FindManyQuery{
		Limit:  input.Limit,
		Offset: input.Offset,
	})
	if err != nil {
		return nil, err
	}
	products := make([]ProductItem, 0, len(productEntities))
	for _, product := range productEntities {
		products = append(products, ProductItem{
			ID:    product.ID(),
			Name:  product.Name(),
			Stock: product.Stock(),
		})
	}
	return products, nil
}
