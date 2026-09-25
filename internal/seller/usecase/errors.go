package usecase

import "mini-shop/internal/pkg/err/bizerr"

var (
	EmailExistsErr     = bizerr.BadRequestErr("EMAIL_EXISTS")
	InvalidEmailErr    = bizerr.BadRequestErr("INVALID_EMAIL")
	InvalidPasswordErr = bizerr.BadRequestErr("INVALID_PASSWORD")
)
