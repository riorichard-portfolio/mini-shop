package entity

import "mini-shop/internal/pkg/err/bizerr"

var (
	InvalidQuantityErr   = bizerr.BadRequestErr("INVALID_QUANTITY")
	UnauthorizedSellerErr = bizerr.ForbiddenErr("UNAUTHORIZED_SELLER")
)
