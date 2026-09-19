package jwtprovider

import "mini-shop/internal/pkg/err/bizerr"

var (
	InvalidCustomerToken = bizerr.UnauthorizedErr("INVALID_CUSTOMER_TOKEN")
	InvalidSigningMethodErr = bizerr.UnauthorizedErr("INVALID_SIGNING_METHOD")
)
