package manager

import "simple-bank/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.NewCtx(), r.infra.NewQuery())
}

func (r *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.NewCtx(), r.infra.NewQuery(), r.infra.NewPool())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
