package fiberhttp

import (
	"strings"

	"github.com/gofiber/fiber/v3"

	"mini-shop/internal/seller/jwtprovider"
)

type Middleware struct {
	jwt *jwtprovider.JWTProvider
}

func NewMiddleware(
	jwt *jwtprovider.JWTProvider,
) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

func (m *Middleware) AuthJWT(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return UnauthorizedAuthErr
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		return UnauthorizedAuthErr
	}
	jwtPayload, err := m.jwt.Verify(tokenStr)
	if err != nil {
		return UnauthorizedAuthErr
	}
	c.Locals("sellerID", jwtPayload.SellerID)
	return c.Next()
}
