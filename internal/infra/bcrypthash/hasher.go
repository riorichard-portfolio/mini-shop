package bcrypthash

import (
	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (bh *BcryptHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash with bcrypt")
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
		return false, errors.Wrap(err, "failed to compare hash with bcrypt")
	}
	return true, nil
}
