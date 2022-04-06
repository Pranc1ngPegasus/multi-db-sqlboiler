package usecase

import (
	"context"

	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/domain/model"
	"github.com/Pranc1ngPegasus/multi-db-sqlboiler/domain/repository"
)

type (
	RegisterUser interface {
	}

	registerUser struct {
		userRepository repository.User
	}
)

func NewRegisterUser(
	userRepository repository.User,
) RegisterUser {
	return &registerUser{
		userRepository: userRepository,
	}
}

type (
	RegisterUserInput struct {
		Name string
	}

	RegisterUserOutput struct {
		User *model.User
		Err  error
	}
)

func (u *registerUser) Do(ctx context.Context, input RegisterUserInput) *RegisterUserOutput {
	original := model.NewUser(
		input.Name,
	)

	user, err := u.userRepository.Create(original)
	if err != nil {
		return &RegisterUserOutput{
			Err: err,
		}
	}

	return &RegisterUserOutput{
		User: user,
		Err:  nil,
	}
}
