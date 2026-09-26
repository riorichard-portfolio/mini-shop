package order

import (
	"mini-shop/internal/order/gormrepo"
	"mini-shop/internal/order/usecase"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo      *gormrepo.GormRepo
	gormTxmanager *gormrepo.TxManager

	usc *usecase.Usecase
}

func (c *Container) Usecase(
	productSvc usecase.ProductSvc,
) *usecase.Usecase {
	if c.usc != nil {
		return c.usc
	}
	c.usc = usecase.NewUsecase(
		c.GormRepo,
		productSvc,
		c.gormTxmanager,
	)
	return c.usc
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	gormTxmanager := gormrepo.NewTxManager(db)
	return &Container{
		GormRepo:      gormRepo,
		gormTxmanager: gormTxmanager,
	}
}
