package gormrepo

import "mini-shop/internal/pkg/err/bizerr"

var(
	CustomerNotFoundErr = bizerr.NotFoundErr("CUSTOMER_NOT_FOUND")
)