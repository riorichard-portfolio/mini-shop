package jwtprovider

import "mini-shop/internal/pkg/err/bizerr"

var (
	InvalidSellerToken = bizerr.UnauthorizedErr("INVALID_SELLER_TOKEN")
	InvalidSigningMethodErr = bizerr.UnauthorizedErr("INVALID_SIGNING_METHOD")
)
