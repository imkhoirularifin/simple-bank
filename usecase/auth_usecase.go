package usecase

import (
	"errors"
	"simple-bank/db/dto"
	"simple-bank/repository"
	"simple-bank/utils/common"
)

type AuthUseCase interface {
	Login(arg dto.LoginParams) (bool, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func (a *authUseCase) Login(arg dto.LoginParams) (bool, error) {
	user, err := a.userRepo.GetUserWithEmail(arg.Email)
	if err != nil {
		return false, errors.New("invalid email or password")
	}

	// compare password
	err = common.ComparePasswordHash(user.Password, arg.Password)
	if err != nil {
		return false, errors.New("invalid email or password")
	}

	return true, nil
}

func NewAuthUseCase(userRepo repository.UserRepository) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}
