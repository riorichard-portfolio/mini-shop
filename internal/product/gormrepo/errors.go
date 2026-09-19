package gormrepo

import "mini-shop/internal/pkg/err/bizerr"

var(
	ProductNotFoundErr = bizerr.NotFoundErr("PRODUCT_NOT_FOUND")
)