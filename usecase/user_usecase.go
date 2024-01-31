package usecase

import (
	"errors"
	db "simple-bank/db/sqlc"
	"simple-bank/repository"
	"simple-bank/utils/common"
)

type UserUseCase interface {
	CreateUser(arg db.CreateUserParams) (db.User, error)
	GetUser(id string) (db.User, error)
	GetUserWithEmail(email string) (db.User, error)
	ListUsers() ([]db.User, error)
	UpdateUser(arg db.UpdateUserParams) (db.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) CreateUser(arg db.CreateUserParams) (db.User, error) {
	// check if username already exist
	_, err := u.userRepo.GetUserWithUsername(arg.Username)
	if err == nil {
		return db.User{}, errors.New("username already exist")
	}

	// check if email contain space
	isContainSpace := common.IsContainSpace(arg.Email)
	if isContainSpace {
		return db.User{}, errors.New("email must not contain spaces")
	}

	// check if email already exist
	_, err = u.userRepo.GetUserWithEmail(arg.Email)
	if err == nil {
		return db.User{}, errors.New("email already exist")
	}

	// hash password
	hashedPassword, err := common.GeneratePasswordHash(arg.Password)
	if err != nil {
		return db.User{}, err
	}
	arg.Password = hashedPassword

	// store user
	user, err := u.userRepo.CreateUser(arg)
	if err != nil {
		return db.User{}, err
	}

	user.Password = ""

	return user, nil
}

func (u *userUseCase) GetUser(id string) (db.User, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return db.User{}, err
	}

	user.Password = ""

	return user, nil
}

func (u *userUseCase) GetUserWithEmail(email string) (db.User, error) {
	user, err := u.userRepo.GetUserWithEmail(email)
	if err != nil {
		return db.User{}, err
	}

	user.Password = ""

	return user, nil
}

func (u *userUseCase) ListUsers() ([]db.User, error) {
	users, err := u.userRepo.ListUsers()
	if err != nil {
		return []db.User{}, err
	}

	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (u *userUseCase) UpdateUser(arg db.UpdateUserParams) (db.User, error) {
	// check if username already exist
	_, err := u.userRepo.GetUserWithUsername(arg.Username)
	if err == nil {
		return db.User{}, errors.New("username already exist")
	}

	// check if email contain space
	isContainSpace := common.IsContainSpace(arg.Email)
	if isContainSpace {
		return db.User{}, errors.New("email must not contain spaces")
	}

	// check if email already exist
	_, err = u.userRepo.GetUserWithEmail(arg.Email)
	if err == nil {
		return db.User{}, errors.New("email already exist")
	}

	// hash password
	if arg.Password != "" {
		hashedPassword, err := common.GeneratePasswordHash(arg.Password)
		if err != nil {
			return db.User{}, err
		}

		arg.Password = hashedPassword
	}

	user, err := u.userRepo.UpdateUser(arg)
	if err != nil {
		return db.User{}, err
	}

	user.Password = ""

	return user, nil
}

func (u *userUseCase) DeleteUser(id string) error {
	// check if user exist
	_, err := u.userRepo.GetUser(id)
	if err != nil {
		return errors.New("user not found")
	}

	err = u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
