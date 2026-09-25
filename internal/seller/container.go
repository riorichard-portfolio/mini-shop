package seller

import (
	"mini-shop/internal/seller/fiberhttp"
	"mini-shop/internal/seller/gormrepo"
	"mini-shop/internal/seller/jwtprovider"
	"mini-shop/internal/seller/usecase"

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