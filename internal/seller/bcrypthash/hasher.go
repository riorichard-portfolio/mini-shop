package bcrypthash

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func (bh *BcryptHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", fmt.Errorf("seller.bcrypthash.Hash: %w", err)
	}
	return string(hashed), nil
}

func (bh *BcryptHasher) Verify(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("seller.bcrypthash.Verify: %w", err)
	}
	return true, nil
}
