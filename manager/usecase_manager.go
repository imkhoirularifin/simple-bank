package manager

import "simple-bank/usecase"

type UsecaseManager interface {
	UserUseCase() usecase.UserUseCase
	TransactionUseCase() usecase.TransactionUseCase
	AuthUseCase() usecase.AuthUseCase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func (u *usecaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repo.TransactionRepo(), u.repo.UserRepo())
}

func (u *usecaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repo.UserRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
