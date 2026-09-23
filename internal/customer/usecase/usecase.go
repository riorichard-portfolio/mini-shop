package usecase

import (
	"context"

	"github.com/google/uuid"

	"mini-shop/internal/customer/entity"
	"mini-shop/internal/customer/dto"
)

type Repo interface {
	SaveNew(ctx context.Context, customer entity.Customer) error
	IsEmailExists(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (entity.Customer, error)
}

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password string, hashedPassword string) (bool, error)
}

type TokenProvider interface {
	Generate(payload dto.TokenPayload) (string, error)
	Verify(token string) (dto.TokenPayload, error)
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

func (u *Usecase) Register(ctx context.Context, input dto.RegisterInput) error {
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
	customer := entity.NewCustomer(
		uuid.NewString(),
		input.Email,
		hashedPassword,
	)
	err = u.repo.SaveNew(ctx, customer)
	return err
}

func (u *Usecase) Login(ctx context.Context, input dto.LoginInput) (string, error) {
	customer, err := u.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}
	isPasswordCorrect, err := u.hasher.Verify(input.Password, customer.HashedPassword())
	if err != nil {
		return "", err
	}
	if !isPasswordCorrect {
		return "", InvalidPasswordErr
	}
	token, err := u.tokenProvider.Generate(dto.TokenPayload{
		CustomerID: customer.ID(),
	})
	return token, err
}
