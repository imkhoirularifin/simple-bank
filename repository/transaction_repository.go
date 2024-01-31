package repository

import (
	"context"
	db "simple-bank/db/sqlc"

	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository interface {
	CreateTransaction(arg db.CreateTransactionParams) (db.Transaction, error)
	GetTransaction(id string) (db.Transaction, error)
}

type transactionRepository struct {
	ctx   context.Context
	query *db.Queries
	pool  *pgxpool.Pool
}

func NewTransactionRepository(ctx context.Context, db *db.Queries, pool *pgxpool.Pool) TransactionRepository {
	return &transactionRepository{
		ctx:   ctx,
		query: db,
		pool:  pool,
	}
}

func (t *transactionRepository) CreateTransaction(arg db.CreateTransactionParams) (db.Transaction, error) {
	tx, err := t.pool.Begin(t.ctx)
	if err != nil {
		return db.Transaction{}, err
	}
	defer tx.Rollback(t.ctx)

	qtx := t.query.WithTx(tx)

	// get both users
	sender, err := qtx.GetUser(t.ctx, arg.FromUserID)
	if err != nil {
		return db.Transaction{}, err
	}

	receiver, err := qtx.GetUser(t.ctx, arg.ToUserID)
	if err != nil {
		return db.Transaction{}, err
	}

	// update balance
	sender.Balance -= arg.Amount
	receiver.Balance += arg.Amount

	// update balance in database
	_, err = qtx.UpdateUserBalance(t.ctx, db.UpdateUserBalanceParams{
		ID:      sender.ID,
		Balance: sender.Balance,
	})
	if err != nil {
		return db.Transaction{}, err
	}

	_, err = qtx.UpdateUserBalance(t.ctx, db.UpdateUserBalanceParams{
		ID:      receiver.ID,
		Balance: receiver.Balance,
	})
	if err != nil {
		return db.Transaction{}, err
	}

	transaction, err := qtx.CreateTransaction(t.ctx, arg)
	if err != nil {
		return db.Transaction{}, err
	}

	if err := tx.Commit(t.ctx); err != nil {
		return db.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepository) GetTransaction(id string) (db.Transaction, error) {
	pgUUID := convert.StringToUUID(id)

	transaction, err := t.query.GetTransaction(t.ctx, pgUUID)
	if err != nil {
		return db.Transaction{}, err
	}

	return transaction, nil
}
