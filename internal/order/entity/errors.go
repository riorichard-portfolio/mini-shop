package entity

import "mini-shop/internal/pkg/err/bizerr"

var (
	InvalidQuantityErr     = bizerr.BadRequestErr("INVALID_QUANTITY")
	InvalidOrderStatusErr  = bizerr.BadRequestErr("INVALID_ORDER_STATUS")
	UnauthorizedSellerErr  = bizerr.ForbiddenErr("UNAUTHORIZED_SELLER")
	NotAllowedProcessErr   = bizerr.BadRequestErr("NOT_ALLOWED_PROCESS")
	StatusChangeInvalidErr = bizerr.BadRequestErr("STATUS_CHANGE_INVALID")
)
