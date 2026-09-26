package main

import (
	"fmt"
	"log"

	"mini-shop/internal/infra/bcrypthash"
	"mini-shop/internal/infra/config/envconfig"
	"mini-shop/internal/infra/fiberhttp"
	"mini-shop/internal/infra/gormdb"
	"mini-shop/internal/infra/validator"

	"mini-shop/internal/customer"
	"mini-shop/internal/order"
	"mini-shop/internal/product"
	"mini-shop/internal/seller"
)

func main() {
	bcryptHasher := bcrypthash.NewBcryptHasher()
	val, err := validator.New()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	appCfg, err := envconfig.NewConfig(".env", val)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	gormDB, err := gormdb.NewDB(appCfg.PostgreConfig)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fiberErrHandler := fiberhttp.NewErrHandler(val.Translator)
	fiberApp := fiberhttp.NewApp(fiberErrHandler, appCfg.AppConfig.TimeoutDur)

	customerCont, err := customer.NewContainer(
		gormDB,
		"jwt_keys/access_token_jwt_private_key.pem",
		"jwt_keys/access_token_jwt_public_key.pem",
	)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	productCont := product.NewContainer(gormDB)

	orderCont := order.NewContainer(gormDB)

	sellerCont, err := seller.NewContainer(
		gormDB,
		"jwt_keys/access_token_jwt_private_key.pem",
		"jwt_keys/access_token_jwt_public_key.pem",
		bcryptHasher,
	)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	customerUsc := customerCont.Usecase(bcryptHasher)
	productUsc := productCont.Usecase()
	orderUsc := orderCont.Usecase(productCont.Service)
	sellerUsc := sellerCont.Usecase(bcryptHasher)

	customer.RegisterAPI(fiberApp, customerUsc)
	product.RegisterAPI(
		fiberApp,
		productUsc,
		sellerCont.Middleware.AuthJWT,
		customerCont.Middleware.AuthJWT,
	)
	order.RegisterAPI(
		fiberApp,
		orderUsc,
		sellerCont.Middleware.AuthJWT,
		customerCont.Middleware.AuthJWT,
	)
	seller.RegisterAPI(fiberApp, sellerUsc)
	addr := fmt.Sprintf(":%d", appCfg.AppConfig.Port)

	if err := fiberApp.Listen(addr); err != nil {
		log.Fatalf("[FATAL] Server failed to start: %v", err)
	}
}
