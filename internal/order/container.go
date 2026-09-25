package order

import (
	"mini-shop/internal/order/gormrepo"
	"mini-shop/internal/order/usecase"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo *gormrepo.GormRepo
	Usecase  *usecase.Usecase
}

func NewContainer(
	db *gorm.DB,
	productSvc usecase.ProductSvc,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	gormTxmanager := gormrepo.NewTxManager(db)
	usc := usecase.NewUsecase(
		gormRepo,
		productSvc,
		gormTxmanager,
	)
	return &Container{
		GormRepo: gormRepo,
		Usecase:  usc,
	}
}
