package product

import (
	"mini-shop/internal/product/gormrepo"
	"mini-shop/internal/product/service"
	"mini-shop/internal/product/usecase"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo *gormrepo.GormRepo
	Service  *service.Service

	usc *usecase.Usecase
}

func (c *Container) Usecase() *usecase.Usecase {
	if c.usc != nil {
		return c.usc
	}
	c.usc = usecase.NewUsecase(c.GormRepo)
	return c.usc
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	svc := service.NewService(gormRepo)

	return &Container{
		GormRepo: gormRepo,
		Service:  svc,
	}
}
