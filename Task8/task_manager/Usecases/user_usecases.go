package Usecases

import (
	"context"
	"errors"
	"task_manager/Domain"
	"task_manager/Repositories"
	"task_manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user *Domain.User) (*Domain.User, error)
	LoginUser(ctx context.Context, username, password string) (string, error)
	PromoteUser(ctx context.Context, id string) error
}

type UserUsecaseImpl struct {
	userRepo      Repositories.UserRepository
	jwtService Infrastructure.JWTService
	passwordService Infrastructure.PasswordService
}

func NewUserUsecase(userRepo Repositories.UserRepository, jwtService Infrastructure.JWTService, passwordService Infrastructure.PasswordService) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo, jwtService: jwtService, passwordService: passwordService}
}

func (uc *UserUsecaseImpl) RegisterUser(ctx context.Context, user *Domain.User) (*Domain.User, error) {
	createdUser, err := uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (uc *UserUsecaseImpl) LoginUser(ctx context.Context, username, password string) (string, error) {
	user, err := uc.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !uc.passwordService.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := uc.jwtService.GenerateToken(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (uc *UserUsecaseImpl) PromoteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	err = uc.userRepo.PromoteUser(ctx, objID)
	if err != nil {
		return err
	}
	return nil
}