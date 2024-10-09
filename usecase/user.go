package usecase

import (
	"context"
	"udemy_slack_app/model"
	"udemy_slack_app/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, user *model.User) (string, error)
	GetByID(ctx context.Context, userID string) (*model.User, error)
	Update(ctx context.Context, user *model.User, userID string) error
	Delete(ctx context.Context, userID string) error
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := uu.ur.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (uu *userUsecase) GetByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := uu.ur.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) Update(ctx context.Context, user *model.User, userID string) error {
	err := uu.ur.Update(ctx, user, userID)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) Delete(ctx context.Context, userID string) error {
	err := uu.ur.Delete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
