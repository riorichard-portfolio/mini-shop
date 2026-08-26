package seller

import (
	"context"

	"github.com/google/uuid"

	"mini-shop/internal/pkg/bizerr"
	"mini-shop/internal/seller/entity"
)

type Repo interface {
	Save(ctx context.Context, seller *entity.Seller) error
	IsEmailExists(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*entity.Seller, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password string, hashedPassword string) (bool, error)
}

type TokenProvider interface {
	Generate(payload *TokenPayload) (string, error)
	Verify(token string) (TokenPayload, error)
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

func (u *Usecase) Register(ctx context.Context, input RegisterInput) (err error) {
	isEmailExists, err := u.repo.IsEmailExists(ctx, input.Email)
	if err != nil {
		return
	}
	if isEmailExists {
		err = bizerr.New("EMAIL_EXISTS")
		return
	}
	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return
	}
	seller := entity.NewSeller(
		uuid.NewString(),
		input.Email,
		hashedPassword,
	)
	err = u.repo.Save(ctx, seller)
	return
}

func (u *Usecase) Login(ctx context.Context, input LoginInput) (token string, err error) {
	seller, err := u.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return
	}
	if seller == nil {
		err = bizerr.New("INVALID_EMAIL")
		return
	}
	isPasswordCorrect, err := u.hasher.Verify(input.Password, seller.HashedPassword())
	if err != nil {
		return
	}
	if !isPasswordCorrect {
		err = bizerr.New("INVALID_PASSWORD")
		return
	}
	token, err = u.tokenProvider.Generate(&TokenPayload{
		SellerID: seller.ID(),
	})
	return
}
