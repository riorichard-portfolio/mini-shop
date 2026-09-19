package service

import "mini-shop/internal/pkg/err/bizerr"

var(
	InsufficientStockErr = bizerr.BadRequestErr("INSUFFICIENT_STOCK")
	InvalidProductErr = bizerr.BadRequestErr("INVALID_PRODUCT")
)