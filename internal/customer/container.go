package customer

import (
	"mini-shop/internal/customer/fiberhttp"
	"mini-shop/internal/customer/gormrepo"
	"mini-shop/internal/customer/jwtprovider"
	"mini-shop/internal/customer/usecase"

	"mini-shop/internal/infra/bcrypthash"

	"gorm.io/gorm"
)

type Container struct {
	GormRepo    *gormrepo.GormRepo
	JWTProvider *jwtprovider.JWTProvider
	Middleware  *fiberhttp.Middleware
	Usecase     *usecase.Usecase
}

func NewContainer(
	db *gorm.DB,
	jwtPrivateKeyPath string,
	jwtPublicKeyPath string,
	bcryptHasher *bcrypthash.BcryptHasher,
) (*Container, error) {
	gormRepo := gormrepo.NewRepo(db)
	jwtProvider, err := jwtprovider.NewProvider(
		jwtPrivateKeyPath,
		jwtPublicKeyPath,
	)
	if err != nil {
		return nil, err
	}
	middleware := fiberhttp.NewMiddleware(
		jwtProvider,
	)
	usc := usecase.NewUsecase(
		gormRepo,
		bcryptHasher,
		jwtProvider,
	)
	return &Container{
		GormRepo:    gormRepo,
		JWTProvider: jwtProvider,
		Middleware:  middleware,
		Usecase:     usc,
	}, nil
}
