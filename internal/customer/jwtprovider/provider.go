package jwtprovider

import (
	"crypto/rsa"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"mini-shop/internal/customer"
)

type JWTProvider struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewProvider(
	privateKeyPath string,
	publicKeyPath string,
) (*JWTProvider, error) {
	privBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file from private key path")
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate private key with jwt")
	}

	pubBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file from public key path")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate public key with jwt")
	}

	return &JWTProvider{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (jp *JWTProvider) Generate(payload customer.TokenPayload) (string, error) {
	claims := jwt.MapClaims{
		"customer_id": payload.CustomerID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(jp.privateKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token with jwt")
	}

	return signedToken, nil
}

func (jp *JWTProvider) Verify(tokenStr string) (customer.TokenPayload, error) {
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid signing method in jwt token")
		}
		return jp.publicKey, nil
	})
	if err != nil {
		return customer.TokenPayload{}, errors.CombineErrors(err, InvalidCustomerToken)
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		customerID, ok := claims["customer_id"].(string)
		if !ok || customerID == "" {
			return customer.TokenPayload{}, InvalidCustomerToken
		}
		return customer.TokenPayload{
			CustomerID: customerID,
		}, nil
	}

	return customer.TokenPayload{}, InvalidCustomerToken
}
