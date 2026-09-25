package product

import (
	"mini-shop/internal/product/gormrepo"
	"mini-shop/internal/product/service"
	"mini-shop/internal/product/usecase"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo *gormrepo.GormRepo
	Usecase  *usecase.Usecase
	Service  *service.Service
}

func NewContainer(
	db *gorm.DB,
) *Container {
	gormRepo := gormrepo.NewRepo(db)
	usc := usecase.NewUsecase(gormRepo)
	svc := service.NewService(gormRepo)

	return &Container{
		GormRepo: gormRepo,
		Usecase:  usc,
		Service:  svc,
	}
}
