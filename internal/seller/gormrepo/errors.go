package gormrepo

import "mini-shop/internal/pkg/err/bizerr"

var(
	SellerNotFoundErr = bizerr.NotFoundErr("SELLER_NOT_FOUND")
)