package customer

import(
	"errors"
)

var(
	EmailExistsErr = errors.New("EMAIL_EXISTS")
	InvalidEmailErr = errors.New("INVALID_EMAIL")
	InvalidPasswordErr = errors.New("INVALID_PASSWORD")
)