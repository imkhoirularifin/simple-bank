package repository

import (
	"context"
	db "simple-bank/db/sqlc"

	"github.com/emicklei/pgtalk/convert"
)

type UserRepository interface {
	CreateUser(arg db.CreateUserParams) (db.User, error)
	GetUser(id string) (db.User, error)
	GetUserWithUsername(username string) (db.User, error)
	GetUserWithEmail(email string) (db.User, error)
	ListUsers() ([]db.User, error)
	UpdateUser(arg db.UpdateUserParams) (db.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	ctx   context.Context
	query *db.Queries
}

func NewUserRepository(ctx context.Context, db *db.Queries) UserRepository {
	return &userRepository{
		ctx:   ctx,
		query: db,
	}
}

func (u *userRepository) CreateUser(arg db.CreateUserParams) (db.User, error) {
	user, err := u.query.CreateUser(u.ctx, arg)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUser(id string) (db.User, error) {
	pgUUID := convert.StringToUUID(id)

	user, err := u.query.GetUser(u.ctx, pgUUID)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserWithUsername(username string) (db.User, error) {
	user, err := u.query.GetUserWithUsername(u.ctx, username)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserWithEmail(email string) (db.User, error) {
	user, err := u.query.GetUserWithEmail(u.ctx, email)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *userRepository) ListUsers() ([]db.User, error) {
	users, err := u.query.ListUsers(u.ctx)
	if err != nil {
		return []db.User{}, err
	}

	return users, nil
}

func (u *userRepository) UpdateUser(arg db.UpdateUserParams) (db.User, error) {
	user, err := u.query.UpdateUser(u.ctx, arg)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (u *userRepository) DeleteUser(id string) error {
	pgUUID := convert.StringToUUID(id)
	err := u.query.DeleteUser(u.ctx, pgUUID)
	if err != nil {
		return err
	}

	return nil
}
