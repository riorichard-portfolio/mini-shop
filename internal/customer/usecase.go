package customer

import (
	"context"

	"github.com/google/uuid"

	"mini-shop/internal/customer/entity"
)

type Repo interface {
	SaveNew(ctx context.Context, customer *entity.Customer) error
	IsEmailExists(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*entity.Customer, error)
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
		err = EmailExistsErr
		return
	}
	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return
	}
	customer := entity.NewCustomer(
		uuid.NewString(),
		input.Email,
		hashedPassword,
	)
	err = u.repo.SaveNew(ctx, customer)
	return
}

func (u *Usecase) Login(ctx context.Context, input LoginInput) (token string, err error) {
	customer, err := u.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return
	}
	if customer == nil {
		err = InvalidEmailErr
		return
	}
	isPasswordCorrect, err := u.hasher.Verify(input.Password, customer.HashedPassword())
	if err != nil {
		return
	}
	if !isPasswordCorrect {
		err = InvalidPasswordErr
		return
	}
	token, err = u.tokenProvider.Generate(&TokenPayload{
		CustomerID: customer.ID(),
	})
	return
}
