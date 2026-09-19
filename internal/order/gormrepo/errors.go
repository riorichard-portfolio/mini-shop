package gormrepo

import "mini-shop/internal/pkg/err/bizerr"

var(
	OrderNotFoundErr = bizerr.NotFoundErr("ORDER_NOT_FOUND")
)