package seller

import (
	"context"

	"github.com/google/uuid"

	"mini-shop/internal/seller/entity"
)

type Repo interface {
	SaveNew(ctx context.Context, seller entity.Seller) error
	IsEmailExists(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (entity.Seller, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password string, hashedPassword string) (bool, error)
}

type TokenProvider interface {
	Generate(payload TokenPayload) (string, error)
	Verify(tokenStr string) (TokenPayload, error)
}

type Usecase struct {
	repo          Repo
	hasher        Hasher
	tokenProvider TokenProvider
}

func NewUsecase(
	repo Repo,
	hasher Hasher,
	tokenProvider TokenProvider,
) *Usecase {
	return &Usecase{
		repo:          repo,
		hasher:        hasher,
		tokenProvider: tokenProvider,
	}
}

func (u *Usecase) Register(ctx context.Context, input RegisterInput) error {
	isEmailExists, err := u.repo.IsEmailExists(ctx, input.Email)
	if err != nil {
		return err
	}
	if isEmailExists {
		return EmailExistsErr
	}
	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	seller := entity.NewSeller(
		uuid.NewString(),
		input.Email,
		hashedPassword,
	)
	err = u.repo.SaveNew(ctx, seller)
	return err
}

func (u *Usecase) Login(ctx context.Context, input LoginInput) (string, error) {
	seller, err := u.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}
	isPasswordCorrect, err := u.hasher.Verify(input.Password, seller.HashedPassword())
	if err != nil {
		return "", err
	}
	if !isPasswordCorrect {
		return "", InvalidPasswordErr
	}
	token, err := u.tokenProvider.Generate(TokenPayload{
		SellerID: seller.ID(),
	})
	return token, err
}
