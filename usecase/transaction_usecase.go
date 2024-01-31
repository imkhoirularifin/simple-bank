package usecase

import (
	"errors"
	db "simple-bank/db/sqlc"
	"simple-bank/repository"

	"github.com/emicklei/pgtalk/convert"
)

type TransactionUseCase interface {
	CreateTransaction(arg db.CreateTransactionParams) (db.Transaction, error)
	GetTransaction(id string) (db.Transaction, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

func (t *transactionUseCase) CreateTransaction(arg db.CreateTransactionParams) (db.Transaction, error) {
	fromUser := convert.UUIDToString(arg.FromUserID)
	toUser := convert.UUIDToString(arg.ToUserID)

	// check if sender and receiver exist
	sender, err := t.userRepo.GetUser(fromUser)
	if err != nil {
		return db.Transaction{}, errors.New("sender not found")
	}

	receiver, err := t.userRepo.GetUser(toUser)
	if err != nil {
		return db.Transaction{}, errors.New("receiver not found")
	}

	// check if sender and receiver are not the same
	if sender.ID == receiver.ID {
		return db.Transaction{}, errors.New("cannot send money to yourself")
	}

	// check if amount param is positive
	if arg.Amount <= 0 {
		return db.Transaction{}, errors.New("amount must be a positive number")
	}

	// check if sender has enough balance
	if sender.Balance < arg.Amount {
		return db.Transaction{}, errors.New("insufficient balance")
	}

	transaction, err := t.transactionRepo.CreateTransaction(arg)
	if err != nil {
		return db.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionUseCase) GetTransaction(id string) (db.Transaction, error) {
	transaction, err := t.transactionRepo.GetTransaction(id)
	if err != nil {
		return db.Transaction{}, err
	}

	return transaction, nil
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}
