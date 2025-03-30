package Usecases

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"task_manager/Domain"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
)

type UserUsecase interface {
	Register(user Domain.User) (*Domain.User, error)
	Login(username, password string) (*Domain.User, error)
	Promote(id primitive.ObjectID) error
}

type UserUsecaseImpl struct {
	userRepo      Repositories.UserRepository
	passwordService Infrastructure.PasswordService
	jwtService    Infrastructure.JWTService
}

func NewUserUsecase(userRepo Repositories.UserRepository, passwordService Infrastructure.PasswordService, jwtService Infrastructure.JWTService) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo, passwordService: passwordService, jwtService: jwtService}
}

func (u *UserUsecaseImpl) Register(user Domain.User) (*Domain.User, error) {
	hashedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return u.userRepo.Create(user)
}

func (u *UserUsecaseImpl) Login(username, password string) (*Domain.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if !u.passwordService.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *UserUsecaseImpl) Promote(id primitive.ObjectID) error {
	return u.userRepo.Promote(id)
}